package main

import (
	"fmt"
)

func main() {
	jianfa(14, 4)
}

func jianfa(m, n int) {
	if m < n {
		fmt.Println(m, n)
	}

	var count int
	for m > n {
		m = m - n
		count++
	}

	fmt.Println(count, n-m)
}

// func weiyi2(m, n int) {
// 	var count int
// 	for m > n {
// 		mul := 1
// 		for mul*n <= m {
//
// 		}
// 	}
// }
