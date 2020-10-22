package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}

		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}

func dChannel() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	dLen := len(arr)
	target := 11
	size := 3

	timer := time.NewTimer(time.Second * 5)
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)

	for i := 0; i < dLen; i += size {
		end := i + size
		if end >= dLen {
			end = dLen - 1
		}
		go doTask(ctx, arr[i:end], target, resultChan)
	}

	select {
	case <-timer.C:
		fmt.Println("time out")
		cancel()
	case <-resultChan:
		fmt.Println("found it")
		cancel()
	}
}

func doTask(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			fmt.Println("cancel by main")
			return
		default:

		}

		if v == target {
			resultChan <- true
			return
		}
	}
}
