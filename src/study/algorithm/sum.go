package main

import (
	"fmt"
	"sort"
)

// 两数之和
func main() {
	arr := []int{5, 12, 6, 3, 9, 2, 1, 7}
	sum := 13
	sumTwo(arr, sum)
	sumThree(arr, sum)
	sumThreeV2(arr, sum)
}

func sumTwo(arr []int, sum int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if v, exists := m[sum-arr[i]]; exists {
			fmt.Println("sum two", sum, v, i)
		} else {
			m[arr[i]] = i
		}
	}
}

func sumThree(arr []int, sum int) {
	for i := 0; i < len(arr); i++ {
		temSum := sum - arr[i]
		m := make(map[int]int)
		for j := i + 1; j < len(arr); j++ {
			if v, exists := m[temSum-arr[j]]; exists {
				fmt.Println("sum three", sum, i, v, j)
			} else {
				m[arr[j]] = j
			}
		}
	}
}

func sumThreeV2(arr []int, sum int) {
	// 排序算法
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		m, n := i+1, len(arr)-1
		tempSum := sum - arr[i]
		for m < n {
			tem := arr[m] + arr[n]
			if tem > tempSum {
				n--
			} else if tem < tempSum {
				m++
			} else {
				fmt.Println(i, m, n)
				n--
			}
		}
	}
}
