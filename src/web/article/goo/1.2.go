package main

import (
	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 从无序链表中移除重复项
func main() {
	head := base.NewLNode()
	base.CreateNodeWithValue(head, []int{1, 3, 1, 1, 5, 5, 5, 7})
	base.PrintNode("before: ", head)

	// SortDelete(head)
	// MapDelete(head)

	base.PrintNode("after: ", head)
}
