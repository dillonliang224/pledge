package main

import (
	"fmt"
	"time"
)

func main() {
	// arr := []int{1, 2, 3}
	// newArr := []*int{}
	// for _, v := range arr {
	// 	arr = append(arr, v)
	// 	v := v
	// 	newArr = append(newArr, &v)
	// }
	// fmt.Println(arr)
	//
	// for _, v := range newArr {
	// 	fmt.Println(*v)
	// }

	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
