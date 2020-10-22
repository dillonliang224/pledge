package main

import (
	"fmt"
)

// 485. 最大连续1的个数
// https://leetcode-cn.com/problems/max-consecutive-ones/
func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 0, 1, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			temp++
		} else {
			temp = 0
		}

		if temp > max {
			max = temp
		}
	}
	return max
}
