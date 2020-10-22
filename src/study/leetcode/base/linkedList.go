package base

// 单链表

type List struct {
	head *Node
	len  int
}

type Node struct {
	val  interface{}
	next *Node
}

// ===== Node =====
func NewNode(value interface{}) *Node {
	node := &Node{Val: value, Next: nil}
	return node
}

func (n *Node) Next() *Node {
	next := n.next
	return next
}

func (n *Node) Val() interface{} {
	if n == nil {
		return nil
	}

	return n.val
}

// ===== List =====
func NewList() *List {
	list := &List{}
	return list
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Len() int {
	return l.len
}

// 单链表，从head插入
func (l *List) Push(value interface{}) {
	node := NewNode(value)
	node.next = l.head
	l.head = node
	l.len++
}

// 单链表，从head移除
func (l *List) Pop() *Node {
	if l.Len() == 0 {
		return nil
	}

	node := l.head
	l.head = l.head.Next()
	l.len--

	return node
}
