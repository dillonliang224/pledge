package base

type Stack interface {
	Size() int
	IsEmpty() bool
	Push(n int)
	Pop() int
	Top() int
}

type Stack1 struct {
	arr  []int
	size int
}

func (s *Stack1) Size() int {
	return s.size
}

func (s *Stack1) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack1) Push(n int) {
	s.arr = append(s.arr, n)
	s.size++
}

func (s *Stack1) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	temp := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--
	return temp
}

func (s *Stack1) Top() int {
	if s.IsEmpty() {
		return -1
	}

	return s.arr[s.size-1]
}

type Stack2 struct {
	list *List
	size int
}

type List struct {
	Data int
	Next *List
}

func (s *Stack2) Size() int {
	return s.size
}

func (s *Stack2) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack2) Push(n int) {
	if s.size == 0 {
		s.size++
		s.list = &List{}
		s.list.Next = &List{n, nil}
	} else {
		s.size++
		head := s.list.Next
		node := &List{n, head}
		s.list.Next = node
	}
}

func (s *Stack2) Pop() int {
	if s.IsEmpty() {
		return -1
	}

	temp := s.list.Next
	s.size--

	s.list.Next = temp.Next

	return temp.Data
}

func (s *Stack2) Top() int {
	if s.IsEmpty() {
		return -1
	}

	temp := s.list.Next
	return temp.Data
}
