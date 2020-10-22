package main

import (
	"fmt"
)

// 46. 全排列
func main() {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	// 标记这个元素是否已经添加到结果集
	visited := make([]bool, len(nums))
	permuteBackTrack(nums, visited, list, &result)

	return result
}

func permuteBackTrack(nums []int, visited []bool, list []int, result *[][]int) {
	// 返回条件：临时结果和输入集合长度一致 才是全排列
	if len(nums) == len(list) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 已经添加过的元素，直接跳过
		if visited[i] {
			continue
		}

		// 添加元素
		list = append(list, nums[i])
		visited[i] = true
		permuteBackTrack(nums, visited, list, result)

		// 移除元素
		visited[i] = false
		list = list[0 : len(list)-1]
	}
}
