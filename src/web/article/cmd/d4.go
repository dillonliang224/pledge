package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			fmt.Printf("%c", i)
			ch <- struct{}{}
			<-ch
		}
	}()

	go func() {
		for i := 0; i < 26; i++ {
			<-ch
			fmt.Printf("%v", i)
			ch <- struct{}{}
		}
	}()

	time.Sleep(5 * time.Millisecond)
}
