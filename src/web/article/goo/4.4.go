package main

import (
	"fmt"
)

// 如何找出数组中丢失的数
func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9}
	fmt.Println(findBySum(arr))
	fmt.Println(findByXOR(arr))
}

func findBySum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	sum := 0
	for i := 1; i <= n+1; i++ {
		sum += i
	}

	for _, num := range arr {
		sum -= num
	}

	return sum
}

func findByXOR(arr []int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	xor := 0
	for i := 1; i <= n+1; i++ {
		xor ^= i
	}

	for _, num := range arr {
		xor ^= num
	}

	return xor
}
