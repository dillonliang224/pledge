package pool

import (
	"errors"
	"sync"
	"sync/atomic"
)

const (
	OPENED = iota
	CLOSED
)

type Pool struct {
	capacity int32
	running  int32
	state    int32
	workers  workerList

	lock        sync.Mutex
	cond        *sync.Cond
	workerCache sync.Pool

	blockingNum int
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = -1
	}

	p := &Pool{
		capacity: int32(size),
	}

	p.workerCache.New = func() interface{} {
		return &Worker{
			pool: p,
			task: make(chan func(), 1),
		}
	}

	p.workers = newWorkerStack(size)
	p.cond = sync.NewCond(&p.lock)
	return p
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

func (p *Pool) Free() int {
	return p.Cap() - p.Running()
}

func (p *Pool) BlockingNum() int {
	return p.blockingNum
}

func (p *Pool) Release() {
	atomic.StoreInt32(&p.state, CLOSED)
	p.lock.Lock()
	p.workers.reset()
	p.lock.Unlock()
}

func (p *Pool) Produce(task func()) error {
	if atomic.LoadInt32(&p.state) == CLOSED {
		return errors.New("pool has closed")
	}

	var w *Worker
	w = p.getWorker()
	w.task <- task
	return nil
}

func (p *Pool) getWorker() (w *Worker) {
	spawnWorker := func() {
		w = p.workerCache.Get().(*Worker)
		w.consume()
	}

	p.lock.Lock()

	w = p.workers.detach()
	if w != nil {
		p.lock.Unlock()
	} else if capacity := p.Cap(); capacity == -1 {
		p.lock.Unlock()
		spawnWorker()
	} else if p.Running() < capacity {
		p.lock.Unlock()
		spawnWorker()
	} else {
	Reentry:
		p.blockingNum++
		p.cond.Wait()
		p.blockingNum--
		if p.Running() == 0 {
			p.lock.Unlock()
			spawnWorker()
			return
		}

		w = p.workers.detach()
		if w == nil {
			goto Reentry
		}

		p.lock.Unlock()
	}

	return
}

func (p *Pool) revertWorker(worker *Worker) bool {
	if capacity := p.Cap(); (capacity > 0 && p.Running() > capacity) || atomic.LoadInt32(&p.state) == CLOSED {
		return false
	}

	p.lock.Lock()
	err := p.workers.insert(worker)
	if err != nil {
		p.lock.Unlock()
		return false
	}

	p.cond.Signal()
	p.lock.Unlock()

	return true
}

func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}
