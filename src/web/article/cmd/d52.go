package main

import (
	"fmt"
)

type X struct {
}

func (x *X) test() {
	fmt.Println(x)
}

func main() {
	var a *X
	a.test()

	x := X{}
	x.test()

	s := "dillon"
	fmt.Println(&[]byte(s)[1])
}
