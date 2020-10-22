package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	fmt.Println(len(a), cap(a))
	a = append(a, 3, 4, 5, 6, 7)
	fmt.Println(len(a), cap(a))
}
