package main

import (
	"fmt"
)

func digui486(i int, j int, nums []int) int {
	if i == j {
		return nums[i]
	}
	pickI := nums[i] - digui486(i+1, j, nums)
	pickJ := nums[j] - digui486(i, j-1, nums)
	return max486(pickI, pickJ)
}

func max486(a, b int) int {
	if a > b {
		return a
	}

	return b
}

var (
	memo = make([]int, 0, 4)
)

func main() {
	nums := []int{1, 5, 233, 7}
	r := digui486(0, len(nums)-1, nums)
	fmt.Println(r)
}
