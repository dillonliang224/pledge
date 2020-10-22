package base

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 链表
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode() *LNode {
	return &LNode{}
}

func CreateNode(node *LNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &LNode{}
		cur.Next.Data = i
		cur = cur.Next
	}
}

func CreateNodeWithValue(node *LNode, value []int) {
	cur := node

	for i := len(value) - 1; i >= 0; i-- {
		newNode := &LNode{}
		newNode.Data = value[i]
		newNode.Next = cur.Next
		cur.Next = newNode
	}

	// for i := 1; i < 8; i++ {
	// 	cur.Next = &LNode{}
	// 	cur.Next.Data = i
	// 	cur = cur.Next
	// }
}

func PrintNode(info string, node *LNode) {
	fmt.Print(info)
	for cur := node.Next; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

func PrintNodeWithoutHeader(info string, node *LNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data, " ")
	}
	fmt.Println()
}

// 链表指向逆序，把当前节点的next改为当前节点的pre
func Reverse(node *LNode) {
	var pre *LNode // 前驱节点
	var cur *LNode // 当前节点
	var next = node.Next

	for next != nil {
		cur = next.Next // 下一个要操作的节点
		next.Next = pre // 当前节点的下一个节点是他的前一个节点
		pre = next      // 那么当前节点变为前驱节点
		next = cur      // 下一个要for循环的节点
	}

	node.Next = pre // 链表的头节点即为最后一个pre
}

// 链表插入逆序，即把下一个节点插入到头部
func InsertReverse(node *LNode) {
	var cur *LNode  // 当前节点
	var next *LNode // 后继节点

	cur = node.Next.Next // 起始从链表第二个节点开始
	node.Next.Next = nil // 设置链表第一个节点为尾节点
	for cur != nil {
		next = cur.Next      // 下一个要for操作的cur， 即当前cur.Next
		cur.Next = node.Next // 把当前节点插入到头部，那么当前节点的下一个节点就是现在第一个节点
		node.Next = cur      // 然后更新头部节点为当前节点
		cur = next           // 获取下一个要操作的节点
	}
}

// 链表递归逆序
func RecursiveReverse(node *LNode) {
	firstNode := node.Next
	newNode := RecursiveReverseChild(firstNode)
	node.Next = newNode
}

func RecursiveReverseChild(node *LNode) *LNode {
	PrintNode("child: ", node)
	if node == nil || node.Next == nil {
		return node
	}

	newHead := RecursiveReverseChild(node.Next)
	PrintNode("newHead: ", newHead)

	node.Next.Next = node
	node.Next = nil

	return newHead
}

// 逆序+顺序输出
func ReverseSortV2(node *LNode) {
	newNode := NewLNode()

	for cur := node.Next; cur != nil; cur = cur.Next {
		head := newNode.Next // 新链表头节点

		nNode := &LNode{}     // 构造新节点
		nNode.Next = head     // 新节点的下一个节点是新链表的头节点
		nNode.Data = cur.Data // 新节点的数据是当前节点的数据

		newNode.Next = nNode // 新链表的头节点即是刚创建的节点
	}

	PrintNode("newNode: ", newNode)
}

// 递归逆序输出
func RecursiveReversePrint(node *LNode) {
	ReversePrint(node.Next)
	fmt.Println()
}

func ReversePrint(node *LNode) {
	if node == nil {
		return
	}

	ReversePrint(node.Next)
	fmt.Print(node.Data, " ")
}

func SortDelete(head *LNode) {
	outCur := head.Next
	var innerCur *LNode
	var innerPre *LNode

	for ; outCur != nil; outCur = outCur.Next {
		for innerCur, innerPre = outCur.Next, outCur; innerCur != nil; {
			if innerCur.Data == outCur.Data {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func MapDelete(head *LNode) {
	m := make(map[interface{}]bool)
	var pre *LNode
	pre = head.Next
	m[pre.Data] = true

	for cur := pre.Next; cur != nil; cur = cur.Next {
		if _, ok := m[cur.Data]; !ok {
			// 标记为已存在
			m[cur.Data] = true
			pre = cur
		} else {
			// 删除元素
			pre.Next = cur.Next
		}
	}
}

// 合并链表
func CombineList(l1, l2 *LNode) *LNode {
	var goo int
	c1 := l1.Next
	c2 := l2.Next
	c3 := NewLNode()
	for c1 != nil && c2 != nil {
		sum := (c1.Data).(int) + (c2.Data).(int) + goo
		temp := 0
		if sum >= 10 {
			goo = 1
			temp = sum - 10
		} else {
			goo = 0
			temp = sum
		}

		n := &LNode{}
		// n.Data = temp
		// n.Next = c3.Next
		c3.Next = n
		c3.Next.Data = temp
		c3 = c3.Next

		c1 = c1.Next
		c2 = c2.Next
	}

	if c1 == nil && c2 != nil {
		for ; c2 != nil; c2 = c2.Next {
			n := &LNode{}

			sum := (c2.Data).(int) + goo
			temp := 0
			if sum >= 10 {
				goo = 1
				temp = sum - 10
			} else {
				goo = 0
				temp = sum
			}

			// n.Data = temp
			// n.Next = c3.Next
			c3.Next = n
			c3.Next.Data = temp
			c3 = c3.Next
			// c3.Next = n
		}
	}

	if c1 != nil && c2 == nil {
		for ; c1 != nil; c1 = c1.Next {
			n := &LNode{}
			sum := (c1.Data).(int) + goo
			temp := 0
			if sum >= 10 {
				goo = 1
				temp = sum - 10
			} else {
				goo = 0
				temp = sum
			}

			// n.Data = temp
			// n.Next = c3.Next
			c3.Next = n
			c3.Next.Data = temp
			c3 = c3.Next
			// c3.Next = n
		}
	}

	if goo == 1 {
		n := &LNode{}
		// n.Data = 1
		// n.Next = c3.Next
		// c3.Next = n
		c3.Next = n
		c3.Next.Data = 1
	}

	// Reverse13(c3)
	return c3
}

// 加法链表
func SumList(l1, l2 *LNode) int {
	s1 := getNum(l1)
	s2 := getNum(l2)
	return s1 + s2
}

func getNum(l *LNode) int {
	var sum int
	count := 0
	for cur := l.Next; cur != nil; cur = cur.Next {
		sum += int(math.Pow10(count)) * (cur.Data).(int)
		count++
	}

	return sum
}

func BuildListFromNum(r int) *LNode {
	n := NewLNode()
	s := strconv.Itoa(r)
	for i, _ := range strings.Split(s, "") {
		t, _ := strconv.Atoi(string(s[i]))
		d := &LNode{}
		d.Data = t
		d.Next = n.Next
		n.Next = d
	}
	return n
}

// 通过快慢指针获取中间节点，并截为两个链表
func GetMidNode(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}

	var slow *LNode  // 慢指针
	var quick *LNode // 快指针
	var slowPre *LNode

	slow = head.Next
	quick = head.Next.Next
	slowPre = head.Next

	for quick != nil && quick.Next != nil {
		slowPre = slow
		slow = slow.Next
		quick = quick.Next.Next
	}

	slowPre.Next = nil

	return slow
}

func ReverseWithoutHeader(node *LNode) *LNode {
	if node == nil || node.Next == nil {
		return node
	}

	var pre *LNode
	var next *LNode

	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre
}

func ReOrder(head *LNode) {
	// 获取中位节点
	midNode := GetMidNode(head.Next)

	cur1 := head.Next

	// 翻转后半部分节点
	PrintNodeWithoutHeader("mid: ", midNode)
	cur2 := ReverseWithoutHeader(midNode)
	PrintNodeWithoutHeader("reverse mid: ", cur2)

	// re order
	var temp *LNode

	for cur1.Next != nil {
		temp = cur1.Next
		cur1.Next = cur2
		cur1 = temp

		temp = cur2.Next
		cur2.Next = cur1
		cur2 = temp
	}

	cur1.Next = cur2
}

func BackListForK(head *LNode, k int) {
	var q *LNode
	var s *LNode
	var n = 1

	q = head.Next
	s = head.Next

	for q.Next != nil {
		if n < k {
			q = q.Next
			n++
		} else {
			q = q.Next
			s = s.Next
		}
	}

	if n < k {
		// 处理链表长度不够k的情况
	}

	temp := s
	s = s.Next
	temp.Next = nil
	PrintNode("head: ", head)
	PrintNodeWithoutHeader("s: ", s)

	q.Next = head.Next
	head.Next = s
	PrintNode("head: ", head)
	PrintNodeWithoutHeader("s: ", s)
}

func IsCircle(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fast = head.Next
	var slow = head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		fmt.Println(slow, fast)
		if slow == fast {
			return slow
		}
	}

	return nil
}

func FanZhuan(head *LNode) *LNode {
	if head == nil || head.Next == nil {
		// TODO
		return head
	}

	pre := head
	cur := head.Next

	for cur != nil && cur.Next != nil {
		next := cur.Next.Next

		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next

		pre = cur
		cur = next
	}

	PrintNode("fanzhuan: ", head)
	return head
}

func FanZhuanK(head *LNode, k int) {
	if head == nil || head.Next == nil {
		return
	}

	var pre = head.Next
	var count = 0
	for cur := head.Next; cur != nil; cur = cur.Next {
		count++
		if count == k {
			temp := cur.Next
			cur.Next = nil
			PrintNodeWithoutHeader("t: ", pre)
			d(pre)
			PrintNodeWithoutHeader("test: ", pre)
			count = 0
			pre = temp
			cur.Next = temp
		} else {
			continue
		}
	}
}

func d(node *LNode) {
	pre := node
	for cur := node.Next; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}
