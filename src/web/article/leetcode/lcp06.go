package main

import (
	"fmt"
)

func main() {
	coins := []int{4, 2, 1}
	count := 0
	for i := 0; i < len(coins); i++ {
		if coins[i]%2 == 0 {
			count += coins[i] / 2
		} else {
			count += coins[i]/2 + 1
		}
	}

	fmt.Println(count)
}
