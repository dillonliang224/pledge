package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(subsets(arr))
	fmt.Println(subsetsV2(arr))
}

func subsets(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	r := make([][]int, 0)
	r = append(r, []int{})
	count := 0

	for count < len(nums) {
		tempR := make([][]int, 0)
		for _, v := range r {
			temp := make([]int, 0)
			temp = append(temp, v...)
			temp = append(temp, nums[count])
			tempR = append(tempR, temp)
		}
		count++

		for _, i := range tempR {
			r = append(r, i)
		}
	}

	return r
}

func subsetsV2(nums []int) (ans [][]int) {
	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
