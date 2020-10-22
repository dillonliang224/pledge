package pool

type Worker struct {
	pool *Pool
	task chan func()
}

func (w *Worker) consume() {
	w.pool.incRunning()
	go func() {
		defer func() {
			w.pool.decRunning()
			w.pool.workerCache.Put(w)

			// 异常处理
			if p := recover(); p != nil {
				panic(p)
			}
		}()

		for f := range w.task {
			if f == nil {
				return
			}
			f()

			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}

type workerList interface {
	len() int
	isEmpty() bool
	insert(worker *Worker) error
	detach() *Worker
	reset()
}

type workerStack struct {
	workers []*Worker
	// expiryWorkers []*Worker
	size int
}

func newWorkerStack(size int) *workerStack {
	return &workerStack{
		workers: make([]*Worker, 0, size),
		size:    size,
	}
}

func (ws *workerStack) len() int {
	return len(ws.workers)
}

func (ws *workerStack) isEmpty() bool {
	return len(ws.workers) == 0
}

func (ws *workerStack) insert(worker *Worker) error {
	ws.workers = append(ws.workers, worker)
	return nil
}

func (ws *workerStack) detach() *Worker {
	l := len(ws.workers)
	if l == 0 {
		return nil
	}

	w := ws.workers[l-1]
	ws.workers = ws.workers[:l-1]
	return w
}

func (ws *workerStack) reset() {
	for i := 0; i < len(ws.workers); i++ {
		ws.workers[i].task <- nil
	}

	ws.workers = ws.workers[:0]
}

func newWorkerList(size int) workerList {
	return newWorkerStack(size)
}
