package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(10)
		}

		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
