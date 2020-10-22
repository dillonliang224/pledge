package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

type Queue struct {
	stackA *base.Stack2 // 入队
	stackB *base.Stack2 // 出队
}

func NewQueue() *Queue {
	return &Queue{
		stackA: &base.Stack2{},
		stackB: &base.Stack2{},
	}
}

func (q *Queue) Push(n int) {
	q.stackA.Push(n)
}

func (q *Queue) Pop() int {
	if q.stackB.IsEmpty() {
		// A 转移到B
		for temp := q.stackA.Pop(); temp != -1; {
			q.stackB.Push(temp)
			temp = q.stackA.Pop()
		}

	}

	return q.stackB.Pop()
}

func main() {
	// 如何用两个栈模拟队列操作
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println("begin...")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println("done...")
}
