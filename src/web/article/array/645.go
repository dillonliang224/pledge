package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findErrorNums([]int{5, 3, 1, 2, 3, 4}))
}

func findErrorNums(nums []int) []int {
	sort.Ints(nums)
	dup := -1
	missing := 1

	for i := 0; i < len(nums); i++ {
		if nums[abs(nums[i])-1] < 0 {
			dup = abs(nums[i])
		} else {
			nums[abs(nums[i])-1] *= -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = nums[i] + 1
		}
	}

	return []int{dup, missing}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
