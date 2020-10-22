package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

func main() {
	// 使用数组/链表实现栈
	// s1 := &base.Stack1{}
	s1 := &base.Stack2{}
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)

	fmt.Println(s1.Size())
	fmt.Println(s1.IsEmpty())

	fmt.Println(s1.Top())

	s1.Pop()
	s1.Pop()
	fmt.Println(s1, s1.IsEmpty(), s1.Size())

	s1.Pop()
	fmt.Println(s1.IsEmpty(), s1.Size())
}
