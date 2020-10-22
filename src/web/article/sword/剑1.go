package main

import (
	"fmt"
)

func main() {
	// LeftRight()
	QuickSlow()
}

func LeftRight() {
	arr := []int{1, 2, 3, 4, 5, 6}

	left := arr[0]
	right := arr[len(arr)-1]
	fmt.Println(left, right)

	for i, j := 0, len(arr)-1; i <= j; {
		if IsJinshu(arr[i]) {
			i++
			continue
		}

		if !IsJinshu(arr[j]) {
			j--
			continue
		}

		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i++
		j--
	}

	fmt.Println(arr)
}

func QuickSlow() {
	arr := []int{1, 2, 3, 4, 5, 6}

	for fast, slow := arr[1], arr[0]; fast < len(arr); {
		if arr[fast]&1 != 0 {
			temp := arr[fast]
			arr[fast] = arr[slow]
			arr[slow] = temp

			slow++
		}

		fast++
	}

	fmt.Println(arr)
}

func IsJinshu(n int) bool {
	return n%2 != 0
}
