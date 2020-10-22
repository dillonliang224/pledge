package main

import (
	"fmt"
	"math"

	"github.com/isdamir/gotype"
)

// 如何求数组中两个元素的最小距离
func main() {
	arr := []int{4, 5, 6, 4, 7, 4, 6, 4, 7, 8, 5, 6, 4, 3, 10, 8}
	get47ByDy(arr, 4, 8)
}

func get47ByVolit(arr []int, a, b int) {
	min := math.MaxInt64
	dist := 0

	for i, v := range arr {
		if v == a {
			for j, v2 := range arr {
				if v2 == b {
					dist = gotype.Abs(i - j)
					if dist < min {
						min = dist
					}
				}
			}
		}
	}

	fmt.Println(min)
}

func get47ByDy(arr []int, a, b int) {
	d1 := -1
	d2 := -1

	min := math.MaxInt64
	for i, v := range arr {
		if v == a {
			d1 = i
		}

		if v == b {
			d2 = i
		}

		min = getMin(min, d1, d2)
	}

	fmt.Println(min)
}

func getMin(min, d1, d2 int) int {
	if d1 == -1 || d2 == -1 {
		return min
	}
	m := 0
	if d1 > 0 && d2 > 0 {
		m = gotype.Abs(d1 - d2)
	}

	if m < min {
		return m
	}
	return min
}
