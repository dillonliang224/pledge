package main

import (
	"fmt"
	"sort"
)

// 78. 子集II
func main() {
	result := subsetsWithDup([]int{1, 2, 2})
	fmt.Println(result)
}

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)

	list := make([]int, 0)
	sort.Ints(nums)
	backtrackII(nums, 0, list, &result)

	return result
}

func backtrackII(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)

	for i := pos; i < len(nums); i++ {
		if i != pos && nums[i] == nums[i-1] {
			continue
		}

		list = append(list, nums[i])
		backtrackII(nums, pos+1, list, result)
		list = list[0 : len(list)-1]
	}
}
