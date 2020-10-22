package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

func main() {
	// 如何实现队列
	// 数组
	// q := &base.Queue1{Arr: make([]int, 0)}

	// 链表
	q := &base.Queue2{}

	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.IsEmpty(), q.Size(), q.GetHead(), q.GetTail())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty(), q.Size(), q.GetHead(), q.GetTail())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty(), q.Size(), q.GetHead(), q.GetTail())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty(), q.Size())

	fmt.Println(q.Pop())
}
