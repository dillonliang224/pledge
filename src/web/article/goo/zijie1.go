package main

import (
	"fmt"
	"math"
)

// 最大化股票收益
func main() {
	arr := []int{100, 80, 120, 130, 70, 60, 100, 125}
	fmt.Println(arr)
	findZijieR1(arr)
	findZijieR2(arr)
	findZiJieR3(arr)
}

func findZijieR1(arr []int) {
	max := math.MinInt64
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			temp := arr[j] - arr[i]
			if temp > max {
				max = temp
			}
		}
	}

	fmt.Println(max)
}

func findZijieR2(arr []int) {
	newArr := make([]int, 0)
	for i, v := range arr {
		if i == 0 {
			newArr = append(newArr, 0)
		} else {
			newArr = append(newArr, v-arr[i-1])
		}
	}

	fmt.Println(newArr)

	maxSum := math.MinInt64
	temp := 0
	for i := 0; i < len(newArr); i++ {
		temp += newArr[i]
		if temp > maxSum {
			maxSum = temp
		}

		if temp < 0 {
			temp = 0
		}
	}

	fmt.Println(maxSum)
}

func findZiJieR3(arr []int) {
	sum := 0
	for i := 1; i < len(arr); i++ {
		temp := arr[i] - arr[i-1]
		if temp > 0 {
			sum += temp
		}
	}

	fmt.Println(sum)
}
