package main

import (
	"fmt"
)

// 90. 子集
func main() {
	result := newbacktrack([]int{1, 2, 3})
	temp := subsets([]int{1, 2, 3})
	fmt.Println(result)
	fmt.Println(temp)
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	list := make([]int, 0)
	backtrack(nums, 0, list, &result)

	return result
}

func newbacktrack(nums []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})
	for _, num := range nums {
		temp := make([][]int, 0)
		for _, before := range result {
			before = append(before, num)
			temp = append(temp, before)
		}

		result = append(result, temp...)
	}

	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
