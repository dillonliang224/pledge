package main

import (
	"fmt"
	"math"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 股票买卖，2次
func main() {
	arr := []int{7, 1, 5, 3, 6, 4}

	fstBuy := math.MinInt64
	fstSell := 0

	secBuy := math.MinInt64
	secSell := 0

	for i := 0; i < len(arr); i++ {
		fstBuy = base.Max(fstBuy, -arr[i])
		fstSell = base.Max(fstSell, fstBuy+arr[i])
		secBuy = base.Max(secBuy, fstSell-arr[i])
		secSell = base.Max(secSell, secBuy+arr[i])
	}

	fmt.Println(secSell)
}
