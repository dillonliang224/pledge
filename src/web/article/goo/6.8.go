package main

import (
	"fmt"
)

func main() {
	count := 1
	var i = 1
	for count <= 1500 {
		i++
		if i%2 == 0 || i%3 == 0 || i%5 == 0 {
			count++
		}
	}

	fmt.Println(i)

}
