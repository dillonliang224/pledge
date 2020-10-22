package main

import (
	"fmt"
)

func main() {
	r := max67(10, 16)
	fmt.Println(r)
}

func max67(a, b int) int {
	if (a-b)&(1<<31) != 0 {
		return b
	}

	return a
}
