package main

import (
	"fmt"
)

func main() {
	r := nthUglyNumber(11)
	fmt.Println(r)
}

// 264

func nthUglyNumber(n int) int {
	if n <= 0 {
		return -1
	}

	arr := make([]int, n)
	arr[0] = 1

	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		next := min(min(arr[p2]*2, arr[p3]*3), arr[p5]*5)
		arr[i] = next

		if next == arr[p2]*2 {
			p2++
		}

		if next == arr[p3]*3 {
			p3++
		}

		if next == arr[p5]*5 {
			p5++
		}

	}

	return arr[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
