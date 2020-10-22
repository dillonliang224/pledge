package base

type Queue interface {
	Pop() int
	Push(n int)
	GetHead() int
	GetTail() int
	Size() int
	IsEmpty() bool
}

type Queue1 struct {
	Arr  []int // 存储值
	head int   // 队列头
	tail int   // 队列尾
}

func (q *Queue1) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue1) Size() int {
	return q.tail - q.head
}

func (q *Queue1) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	temp := q.Arr[q.head]

	q.tail--
	q.Arr = q.Arr[1:]
	return temp
}

func (q *Queue1) Push(n int) {
	q.Arr = append(q.Arr, n)
	q.tail++
}

func (q *Queue1) GetHead() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Arr[q.head]
}

func (q *Queue1) GetTail() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Arr[q.tail-1]
}

type Queue2 struct {
	Head *List
	Tail *List
	size int
}

func (q *Queue2) IsEmpty() bool {
	return q.Head == nil
}

func (q *Queue2) Size() int {
	return q.size
}

func (q *Queue2) Pop() int {
	if q.IsEmpty() {
		return -1
	}

	head := q.Head
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}

	q.size--

	return head.Data
}

func (q *Queue2) Push(n int) {
	node := &List{n, nil}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
	q.size++
}

func (q *Queue2) GetHead() int {
	return q.Head.Data
}

func (q *Queue2) GetTail() int {
	return q.Tail.Data
}
