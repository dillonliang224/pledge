package main

import "fmt"

type MyInterface interface {
	Print()
}

type MyStruct struct {
}

func (ms MyStruct) Print() {

}

func main() {
	a := 1
	b := "str"
	c := MyStruct{}
	var i1 interface{} = a
	var i2 interface{} = b
	var i3 MyInterface = c
	var i4 interface{} = i3
	var i5 = i4.(MyInterface)
	fmt.Println(i1, i2, i3, i4, i5)
}
