package main

import (
	"fmt"
	"time"
)

func main() {
	// go build
	// GODEBUG=schedtrace=1000 ./trace2
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello GMP ", i)
	}
}
