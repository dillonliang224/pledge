package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]int, 5, 9)
	fmt.Println(s)

	Len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s))
	Cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s))
}
