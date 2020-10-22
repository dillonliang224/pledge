package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 股票买卖，不限制次数
func main() {
	arr := []int{7, 1, 5, 3, 6, 4}

	buy := -arr[0]
	sel := 0

	for i := 1; i < len(arr); i++ {
		buy = base.Max(buy, sel-arr[i])
		sel = base.Max(sel, buy+arr[i])
	}

	fmt.Println(sel)
}
