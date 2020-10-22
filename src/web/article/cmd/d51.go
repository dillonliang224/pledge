package main

import (
	"fmt"
)

func main() {
	var f1 = func() {}
	var f2 = func() {}

	if f1 != f2 {
		fmt.Println(11)
	}
}
