package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	// GiB // 1073741824
	// TiB // 1099511627776             (超过了int32的范围)
	// PiB // 1125899906842624
	// EiB // 1152921504606846976
	// ZiB // 1180591620717411303424    (超过了int64的范围)
	// YiB // 1208925819614629174706176
)

var curMem uint64

// https://github.com/panjf2000/ants/blob/master/ants_test.go

func TestPoolWork(t *testing.T) {
	var wg sync.WaitGroup
	p := NewPool(100)
	defer p.Release()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		_ = p.Produce(func() {
			time.Sleep(time.Second)
			fmt.Println("Hello world")
			wg.Done()
		})
	}

	wg.Wait()
	t.Logf("pool, running workers number:%d", p.Running())
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	curMem = mem.TotalAlloc/MiB - curMem
	t.Logf("memory usage:%d MB", curMem)
}
