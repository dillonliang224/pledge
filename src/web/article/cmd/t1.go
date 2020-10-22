package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	A uint32
	B uint64
	C uint64
	D uint32
	E struct{}
}

func main() {
	fmt.Println(unsafe.Offsetof(S{}.E))
	fmt.Println(unsafe.Sizeof(S{}.E))
	fmt.Println(unsafe.Sizeof(S{}))
}
