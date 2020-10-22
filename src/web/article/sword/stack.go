package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

type Queue struct {
	// 出队
	Head *base.Stack2
	// 入队
	Tail *base.Stack2
}

func (q *Queue) Empty() bool {
	return q.Head.IsEmpty() && q.Tail.IsEmpty()
}

func (q *Queue) Push(x int) {
	q.Tail.Push(x)
}

func (q *Queue) Pop() int {
	if q.Head.Top() != -1 {
		return q.Head.Pop()
	}

	if q.Tail.Pop() == -1 {
		return -1
	}

	for node := q.Tail.Pop(); node != -1; node = q.Tail.Pop() {
		q.Head.Push(node)
	}

	return q.Head.Pop()
}

func (q *Queue) Peek() int {
	if q.Head.Top() != -1 {
		return q.Head.Top()
	}

	if q.Tail.Top() == -1 {
		return -1
	}

	for node := q.Tail.Pop(); node != -1; node = q.Tail.Pop() {
		q.Head.Push(node)
	}

	return q.Head.Top()
}

func main() {
	q := &Queue{Head: &base.Stack2{}, Tail: &base.Stack2{}}
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek(), q.Empty())
	q.Push(3)
	fmt.Println(q.Pop(), q.Peek(), q.Empty())
	q.Push(4)
	fmt.Println(q.Pop(), q.Peek(), q.Empty())
	fmt.Println(q.Pop(), q.Peek(), q.Empty())
	fmt.Println(q.Pop(), q.Peek(), q.Empty())
	fmt.Println(q.Pop(), q.Peek(), q.Empty())
}
