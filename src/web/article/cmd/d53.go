package main

import "fmt"

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() T {
	return T{}
}

func main() {
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}
