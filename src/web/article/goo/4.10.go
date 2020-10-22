package main

import (
	"fmt"

	"github.com/isdamir/gotype"
)

// 如何求数组连续最大和
func main() {
	arr := []int{1, -2, 4, 8, -4, 7, -1, -5}
	// find410By1(arr)
	find410ByD(arr)
}

func find410By1(arr []int) {
	maxSum := 0

	for i := 0; i < len(arr); i++ {
		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	fmt.Println("暴力求解法： ", maxSum)
}

// TODO
func find410ByD(arr []int) {
	nAll := arr[0]
	nEnd := arr[0]

	for _, v := range arr {
		nEnd = gotype.Max(nEnd+v, v)
		nAll = gotype.Max(nEnd, nAll)
	}

	fmt.Println("动态规划: ", nAll)
}
