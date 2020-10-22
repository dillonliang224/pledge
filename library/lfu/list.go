package lfu

// 节点
type Node struct {
	key      int
	value    int
	frequent int   // 访问次数
	pre      *Node // 前驱指针
	next     *Node // 后驱指针
}

// 双向链表
type ListNode struct {
	head *Node
	tail *Node
	size int // 尺寸
}

func (l *ListNode) addNode(node *Node) {
	head := l.head

	node.next = head.next
	head.next.pre = node

	node.pre = head
	head.next = node
}

func (l *ListNode) removeNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
