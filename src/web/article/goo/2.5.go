package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

type MinStack struct {
	eles    *base.Stack2
	minEles *base.Stack2
}

func NewMinStack() *MinStack {
	return &MinStack{
		eles:    &base.Stack2{},
		minEles: &base.Stack2{},
	}
}

func (m *MinStack) Push(n int) {
	if m.minEles.IsEmpty() {
		m.minEles.Push(n)
	} else {
		min := m.minEles.Top()
		if n < min {
			m.minEles.Push(n)
		}

	}

	m.eles.Push(n)
}

func (m *MinStack) Pop() int {
	if m.minEles.IsEmpty() {
		return -1
	} else {
		min := m.minEles.Top()
		temp := m.eles.Pop()
		if min == temp {
			m.minEles.Pop()
		}

		return temp
	}
}

func (m *MinStack) Min() int {
	return m.minEles.Top()
}

// 如何用O(1)的时间复杂度求栈中最小元素
func main() {
	// 空间换时间，两个栈
	m := NewMinStack()
	m.Push(5)
	fmt.Println("min: ", m.Min())
	m.Push(6)
	fmt.Println("min: ", m.Min())
	m.Push(1)
	fmt.Println("min: ", m.Min())
	m.Push(2)
	fmt.Println("min: ", m.Min())
	m.Pop()
	fmt.Println("min: ", m.Min())
	m.Pop()
	fmt.Println("min: ", m.Min())
}
