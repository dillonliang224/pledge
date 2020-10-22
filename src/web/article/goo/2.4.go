package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

func main() {
	push := []int{1, 2, 3, 4, 5}
	pop := []int{3, 5, 2, 4, 1}

	stack := &base.Stack2{}
	index := 0
	temp := 0
	for i := 0; i < len(push); i++ {
		stack.Push(push[i])
		for index <= len(pop)-1 && stack.Top() == pop[index] {
			temp = stack.Pop()
			fmt.Println(temp)
			index++
		}
	}

	if index == len(pop) && pop[index-1] == temp {
		fmt.Println("done...")
	} else {
		fmt.Println("panic: ", index)
	}
}
