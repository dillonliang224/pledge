package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 有环
func main() {
	head := base.NewLNode()
	base.CreateNodeWithValue(head, []int{1, 2, 3, 4, 5, 6, 7, 3})
	base.PrintNode("before: ", head)

	// 如何构造有环链表
	r := base.IsCircle(head)
	fmt.Println("list has circle: ", r)
}
