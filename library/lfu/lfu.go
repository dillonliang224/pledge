package lfu

/**
LFU: Least Frequently Used，缓存满的时候，删除缓存里使用次数最少的元素，然后放入新元素，如果使用频率一样，删除缓存最久的元素
*/

// LFUCache结构：包含capacity容量, size当前容量, minFrequent当前最少访问频次, cacheMap缓存哈希表, frequentMap频次哈希表
// minFrequent当前最少访问频次：
// 1. 插入一个新节点时，之前肯定没访问过，minFrequent = 1
// 2. put和get时，如果移除后双向链表节点个数为0，且恰好是最小访问链表, minFrequent++

// 非并发安全的
type LFUCache struct {
	cap         int
	size        int
	minFrequent int
	cacheMap    map[int]*Node
	frequentMap map[int]*ListNode
}

func New(capacity int) *LFUCache {
	return &LFUCache{
		cap:         capacity,
		size:        0,
		minFrequent: 0,
		cacheMap:    make(map[int]*Node),
		frequentMap: make(map[int]*ListNode),
	}
}

func (lfu *LFUCache) Get(key int) int {
	if node, exist := lfu.cacheMap[key]; exist {
		lfu.triggerVisit(node)
		return node.value
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.cap == 0 {
		return
	}

	if node, exist := lfu.cacheMap[key]; exist {
		lfu.triggerVisit(node)
		lfu.cacheMap[key].value = value
	} else {
		newNode := &Node{key, value, 1, nil, nil}
		if lfu.size < lfu.cap {
			lfu.size++
		} else {
			lfu.doLFU()
		}

		lfu.add(newNode)
		lfu.minFrequent = 1
		lfu.cacheMap[key] = newNode
	}
}

// LFUCache辅助函数：获取一个key和修改一个key都会增加对应key的访问频次，可以独立为一个方法，完成如下任务：
// 1. 将对应node从频次列表中移出
// 2. 维护minFrequent
// 3. 该节点访问频次++，移动进下一个访问频次链表
func (lfu *LFUCache) triggerVisit(node *Node) {
	lfu.remove(node)
	if node.frequent == lfu.minFrequent && lfu.frequentMap[node.frequent].size == 0 {
		lfu.minFrequent++
	}
	node.frequent++
	lfu.add(node)
}

// LFUCache辅助函数：将节点添加进对应的频次双向链表，没有则创建
func (lfu *LFUCache) add(node *Node) {
	if listNode, exist := lfu.frequentMap[node.frequent]; exist {
		listNode.addNode(node)
		listNode.size++
	} else {
		listNode := &ListNode{&Node{}, &Node{}, 0}
		listNode.head.next = listNode.tail
		listNode.tail.pre = listNode.head
		listNode.addNode(node)
		listNode.size++
		lfu.cacheMap[node.key] = node
		lfu.frequentMap[node.frequent] = listNode
	}
}

func (lfu *LFUCache) remove(node *Node) {
	lfu.frequentMap[node.frequent].removeNode(node)
	lfu.frequentMap[node.frequent].size--
}

// 执行LFU，删除最近最少使用的一个node
func (lfu *LFUCache) doLFU() {
	listNode := lfu.frequentMap[lfu.minFrequent]
	delete(lfu.cacheMap, listNode.tail.pre.key)
	listNode.removeNode(listNode.tail.pre)
	listNode.size--
}
