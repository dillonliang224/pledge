package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, *v)
	}

	s2 := []int{}
	if s2 == nil {
		fmt.Println(111)
	} else {
		fmt.Println(22)
	}
}
