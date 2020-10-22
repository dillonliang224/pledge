package main

import (
	"fmt"
)

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 3}
	arr := []int{1, 3, 4, 2, 5, 3}
	r := printWithHash(arr)
	fmt.Println(r)

	r5 := printWithCircle(arr)
	fmt.Println(r5)

	r2 := printWithSum(arr)
	fmt.Println(r2)

	r3 := printWithXOR(arr)
	fmt.Println(r3)

	r4 := printWithMark(arr)
	fmt.Println(r4)
}

func printWithHash(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	m := make(map[int]bool)
	for _, v := range arr {
		if m[v] {
			return v
		} else {
			m[v] = true
		}
	}

	return -1
}

func printWithSum(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	for i := 1; i <= len(arr)-1; i++ {
		sum -= i
	}

	return sum
}

func printWithXOR(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	result := 0
	for _, v := range arr {
		result ^= v
	}

	for i := 1; i < len(arr); i++ {
		result ^= i
	}

	return result
}

func printWithMark(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	index := 0

	for {
		if arr[index] < 0 {
			break
		}

		if arr[index] > 0 {
			temp := arr[index]
			arr[index] *= -1
			index = temp
		}

	}

	return index
}

func printWithCircle(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	slow := 0
	fast := 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[arr[fast]]
		slow = arr[slow]
	}

	fast = 0
	for ok := true; ok; ok = slow != fast {
		fast = arr[fast]
		slow = arr[slow]
	}

	return slow
}
