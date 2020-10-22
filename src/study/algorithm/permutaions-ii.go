package main

import (
	"fmt"
	"sort"
)

func main() {
	result := permuteUnique([]int{1, 2, 2, 3})
	fmt.Println(result, len(result))
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	permuteBackTrackII(nums, visited, list, &result)

	return result
}

func permuteBackTrackII(nums []int, visited []bool, list []int, result *[][]int) {
	if len(nums) == len(list) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		if i != 0 && nums[i] == nums[i-1] && visited[i-1] {
			continue
		}

		list = append(list, nums[i])
		visited[i] = true

		permuteBackTrackII(nums, visited, list, result)

		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
