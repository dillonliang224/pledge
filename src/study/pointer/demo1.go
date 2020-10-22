package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	a byte
	c int8
	b int32
}

func main() {
	hello := []byte{104, 101, 108, 108, 111}
	p := unsafe.Pointer(&hello)
	h := (*string)(p)
	fmt.Println(h, *h)

	t := T{}
	fmt.Printf("t size : %d, aligh = %d\n", unsafe.Sizeof(t), unsafe.Alignof(t))
}
