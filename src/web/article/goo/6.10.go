package main

import (
	"fmt"
)

func main() {
	fmt.Println(weiyii610(7))
	fmt.Println(weiyii610(8))
}

func weiyi610(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}

		n >>= 1
	}

	return count
}

func weiyii610(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}
