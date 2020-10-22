package main

import (
	"fmt"
)

// 如何找出数组中出现奇数次的数
func main() {
	arr := []int{3, 5, 6, 6, 5, 7, 2, 2}
	findJiShuByMap(arr)

	findJiShuByXOR(arr)
}

func findJiShuByMap(arr []int) {
	m := make(map[int]bool)

	for _, v := range arr {
		if _, ok := m[v]; ok {
			m[v] = false
		} else {
			m[v] = true
		}
	}

	for k, v := range m {
		if v {
			fmt.Println(k)
		}
	}
}

func findJiShuByXOR(arr []int) {
	r := 0
	for _, v := range arr {
		r ^= v
	}

	temp := r

	position := uint(0)
	for i := r; i&1 == 0; i = i >> 1 {
		position++
	}

	// fmt.Println(position)

	for _, v := range arr {
		if (v>>position)&1 == 1 {
			r ^= v
		}
	}

	fmt.Println(r)
	fmt.Println(r ^ temp)
}
