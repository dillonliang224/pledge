package main

import (
	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 如何找出单链表中的倒数第k个元素
func main() {
	head := base.NewLNode()
	base.CreateNodeWithValue(head, []int{1, 2, 3, 4, 5, 6, 7})
	base.PrintNode("before: ", head)

	base.BackListForK(head, 4)
}
