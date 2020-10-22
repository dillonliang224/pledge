package main

import (
	"fmt"
)

func main() {
	coins := []int{2, 5, 7}
	r := coinChange(coins, 27)
	fmt.Println(r)

	r2 := coinChange2(coins, 27)
	fmt.Println(r2)

	r3 := coinChange3(coins, 27)
	fmt.Println(r3)
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1

		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}

	fmt.Println(dp)
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}

	dp[0] = 0

	// dp[i] = min(dp[i-2], dp[i-5], dp[i-7]) + 1
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 {
				if dp[i] == -1 || dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	fmt.Println(dp)
	return dp[amount]
}

func coinChange3(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	// 动态规划的方程
	// dp[i]=min(dp[i-x],dp[i-y],dp[i-z])+1  x y z 为金额面值
	for i := 1; i <= amount; i++ {
		// 剩余的
		for j := 0; j <= len(coins)-1; j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				// 对每个面值，dp 没有初始化过或者有更小的值，这里取最小的值
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}
