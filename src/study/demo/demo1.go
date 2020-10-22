package main

import (
	"bytes"
	"fmt"
)

func main() {
	c := "ab"
	b := bytes.NewBufferString(c)
	fmt.Println(b.Cap(), b.Len())

	contents := "ab"
	buffer1 := bytes.NewBufferString(contents)
	fmt.Printf("The capacity of new buffer with contents %q: %d\n",
		contents, buffer1.Cap()) // 内容容器的容量为：8。
	unreadBytes := buffer1.Bytes()
	fmt.Printf("The unread bytes of the buffer: %v\n", unreadBytes) // 未读内容为：[97 98]。

	size := uintptr(2)

	a := uintptr(1 << 13)
	t := (size + a - 1) &^ (a - 1)
	fmt.Println(t)
}
