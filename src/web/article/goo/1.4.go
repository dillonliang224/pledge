package main

import (
	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 如何对链表进行重新排序
func main() {
	head := base.NewLNode()
	base.CreateNodeWithValue(head, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	base.PrintNode("before: ", head)

	base.ReOrder(head)

	base.PrintNode("after: ", head)
}
