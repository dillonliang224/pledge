package main

import (
	"fmt"
	"math"

	"github.com/isdamir/gotype"
)

// 如何求数组中绝对值最小的数
func main() {
	arr := []int{-10, -5, -2, 7, 15, 50}
	find49ByOrder(arr)
	find49ByErFen(arr)
}

func find49ByErFen(arr []int) {
	begin := 0
	end := len(arr)
	if arr[0] >= 0 {
		fmt.Println("1111")
	} else if arr[end-1] <= 0 {
		fmt.Println(2222)
	} else {
		mid := 0
		for true {
			mid = begin + (end-begin)/2
			if arr[mid] == 0 {
				fmt.Println(3333)
			} else if arr[mid] > 0 {
				if arr[mid-1] > 0 {
					end = mid - 1
				} else if arr[mid-1] == 0 {
					fmt.Println(0000)
				} else {
					break
				}
			} else {
				if arr[mid+1] == 0 {
					fmt.Println(000)
				} else if arr[mid+1] < 0 {
					begin = mid + 1
				} else {
					break
				}
			}
		}

		if arr[mid] > 0 {
			if arr[mid] >= gotype.Abs(arr[mid-1]) {
				fmt.Println(gotype.Abs(arr[mid-1]))
			} else {
				fmt.Println(arr[mid])
			}
		} else {
			if gotype.Abs(arr[mid]) < arr[mid+1] {
				fmt.Println(gotype.Abs(arr[mid]))
			} else {
				fmt.Println(arr[mid+1])
			}
		}

	}

}

func find49ByOrder(arr []int) {
	min := math.MaxInt64
	for _, v := range arr {
		if gotype.Abs(v) < min {
			min = gotype.Abs(v)
		}
	}

	fmt.Println("排序： ", min)
}
