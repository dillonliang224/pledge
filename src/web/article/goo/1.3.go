package main

import (
	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 计算两个链表所代表的数之和
func main() {
	l1 := base.NewLNode()
	base.CreateNodeWithValue(l1, []int{3, 1, 5})

	l2 := base.NewLNode()
	base.CreateNodeWithValue(l2, []int{5, 9, 4})

	// l3 := buildList(SumList(l1, l2))
	// base.PrintNode("after: ", l3)

	l3 := base.CombineList(l1, l2)
	base.PrintNode("after: ", l3)
}
