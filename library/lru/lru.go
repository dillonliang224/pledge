package lru

/**
  LRU: Least Recently Used，缓存满的时候，删除缓存里最久未使用的数据，然后放入新元素
  m 代码lru里的元素，用map方面查询
  要实现lru，用双端列表head/tail
*/
type LRUCache struct {
	m   map[int]*ListNode
	cap int

	head *ListNode
	tail *ListNode
}

func New(capacity int) *LRUCache {
	head := &ListNode{0, 0, nil, nil}
	tail := &ListNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return &LRUCache{
		m:    make(map[int]*ListNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

// get 会把值放到列表head
func (lru *LRUCache) Get(key int) int {
	cache := lru.m
	if v, exist := cache[key]; exist {
		lru.RemoveNode(v)
		lru.AddNode(v)
		return v.val
	} else {
		return -1
	}
}

// put 新增一个放到列表head，如果列表已满，把最后一个删除
func (lru *LRUCache) Put(key int, value int) {
	cache := lru.m

	// 值存在，更新值并移动到head
	if v, exist := cache[key]; exist {
		v.val = value
		lru.RemoveNode(v)
		lru.AddNode(v)
	} else {
		if len(cache) == lru.cap {
			// 删除最后的元素
			delete(cache, lru.tail.pre.key)
			lru.RemoveNode(lru.tail.pre)
		}

		newNode := &ListNode{key, value, nil, nil}
		lru.AddNode(v)
		cache[key] = newNode
	}
}

// A->B
// C -> A -> B
// 新增C节点，
func (lru *LRUCache) AddNode(node *ListNode) {
	// 获取head
	head := lru.head

	// C节点的next节点是原head节点的next（A）,同样，原head节点的next（A）节点的pre节点将会变为C
	node.next = head.next
	head.next.pre = node

	// head节点的next就是C
	// C节点的pre节点就是head
	node.pre = head
	head.next = node
}

// A -> B -> C
// A -> C
// 删除B节点，那么A节点的next=C, C节点的pre=A
func (lru *LRUCache) RemoveNode(node *ListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
