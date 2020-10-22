package main

import (
	"fmt"
)

func main() {
	// sum := 0
	r := tiao(7)
	fmt.Println(r)

	r2 := ddp(7)
	fmt.Println(r2)

}

func tiao(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return tiao(n-1) + tiao(n-2)
}

func ddp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	// sum := 0
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n]
}
