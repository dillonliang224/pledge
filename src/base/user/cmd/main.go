package main

import (
	"fmt"
	"unicode/utf8"
)

func cacl(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type Student struct {
	name string
}

func main() {
	a := 1
	b := 2
	defer cacl("1", a, cacl("10", a, b))

	a = 0
	defer cacl("2", a, cacl("20", a, b))
	b = 1

	fmt.Println(len("你好"))
	fmt.Println(len([]byte("你好")))
	fmt.Println(utf8.RuneCountInString("你好"))

	m := map[string]Student{"people": {"zhoujielun"}}
	fmt.Println(m["people"], m["people"].name)
	// m["people"].name = "wuyanzu"
}
