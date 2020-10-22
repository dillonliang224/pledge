package main

import (
	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

// 合并两个有序链表
func main() {
	head1 := base.NewLNode()
	base.CreateNodeWithValue(head1, []int{1, 3, 5, 6})

	head2 := base.NewLNode()
	base.CreateNodeWithValue(head2, []int{2, 4, 7, 8, 9})

	var h *base.LNode // 头节点
	var c *base.LNode // 新链表的尾节点

	c1, c2 := head1.Next, head2.Next
	if (c1.Data).(int) > (c2.Data).(int) {
		c = c2
		c2 = c2.Next
	} else {
		c = c1
		c1 = c1.Next
	}

	h = base.NewLNode()
	h.Next = c

	for c1 != nil && c2 != nil {
		if (c1.Data).(int) > (c2.Data).(int) {
			c.Next = c2
			c = c2
			c2 = c2.Next
		} else {
			c.Next = c1
			c = c1
			c1 = c1.Next
		}
	}

	if c1 != nil {
		c.Next = c1
	}

	if c2 != nil {
		c.Next = c2
	}

	base.PrintNode("after: ", h)
}
