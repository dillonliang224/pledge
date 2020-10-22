package main

import (
	"fmt"
)

func main() {
	fmt.Println(yu(5), yu(8))
	fmt.Println(weiyi(5), weiyi(8))
}

func yu(n int) bool {
	return n&(n-1) == 0
}

func weiyi(n int) bool {
	for i := 1; i <= n; {
		if i == n {
			return true
		}

		i <<= 1
	}
	return false
}
