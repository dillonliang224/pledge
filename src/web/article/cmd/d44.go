package main

import (
	"fmt"
)

type info struct {
	result int
}

func work() (int, error) {
	return 11, nil
}

func main() {
	var data info
	var err error
	data.result, err = work()
	fmt.Println(data, err)
}
