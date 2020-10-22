package lru

type ListNode struct {
	key  int
	val  int
	pre  *ListNode
	next *ListNode
}
