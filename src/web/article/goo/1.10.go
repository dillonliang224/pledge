package main

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 给定单链表中某个节点指针，然后删除它
func main() {
	a := 64
	r := a & (a - 1)
	fmt.Println(r)

	c := min(5, 3)
	fmt.Println(c)

	head := base.NewLNode()
	retNode := createNode(head, 5)
	removeNode(retNode)
	base.PrintNode("before: ", head)
}

func removeNode(retNode *base.LNode) {
	if retNode.Next == nil {
		return
	}

	next := retNode.Next.Next
	retNode.Data = retNode.Next.Data
	retNode.Next = next
}

func createNode(node *base.LNode, n int) (retNode *base.LNode) {
	cur := node
	for i := 1; i < 8; i++ {
		cur.Next = &base.LNode{}
		cur.Next.Data = i
		cur = cur.Next

		if n == i {
			retNode = cur
		}
	}

	return
}

func min(a, b int) int {
	return copy(make([]int, a), make([]int, b))
}
