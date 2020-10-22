package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
	fmt.Println(findPoisonedDuration([]int{1, 3, 5, 7, 9, 11, 13, 15}, 1))
	fmt.Println(findPoisonedDuration([]int{1, 2, 5, 7, 9, 11, 13, 15}, 0))
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	du := 0
	duOver := 0
	for i := 0; i < len(timeSeries); i++ {
		if timeSeries[i] >= duOver {
			du += duration
			duOver = timeSeries[i] + duration
		} else {
			temp := timeSeries[i] - timeSeries[i-1]
			du += temp
			duOver += temp
		}
	}

	return du
}
