package snowflake

import (
	"fmt"
	"testing"
)

func TestNewSnowflakeWorker(t *testing.T) {
	worker, err := NewSnowflakeWorker(12)
	if err != nil {
		t.Fatal(err)
	}

	ch := make(chan int64)
	defer close(ch)
	count := 10000
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GenerateId()
			ch <- id
		}()
	}

	m := make(map[int64]int)
	for i := 0; i < count; i++ {
		id := <-ch
		_, ok := m[id]
		if ok {
			t.Error("id is not unique")
			return
		}

		m[id] = 1
	}

	fmt.Println("done...")
}
