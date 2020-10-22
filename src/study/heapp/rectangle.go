package main

import (
	"container/heap"
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func (rec *Rectangle) Area() int {
	return rec.width * rec.height
}

type RectHeap []Rectangle

func (rect RectHeap) Len() int {
	return len(rect)
}

// sort.Interface
func (rect RectHeap) Swap(i, j int) {
	rect[i], rect[j] = rect[j], rect[i]
}

func (rect RectHeap) Less(i, j int) bool {
	return rect[i].Area() < rect[j].Area()
}

func (rect *RectHeap) Push(h interface{}) {
	*rect = append(*rect, h.(Rectangle))
}

func (rect *RectHeap) Pop() (x interface{}) {
	n := len(*rect)
	x = (*rect)[n-1]
	*rect = (*rect)[:n-1]
	return x
}

func main() {
	hp := &RectHeap{}
	for i := 2; i < 6; i++ {
		*hp = append(*hp, Rectangle{i, i})
	}

	fmt.Println(hp)

	heap.Init(hp)
	heap.Push(hp, Rectangle{100, 10})
	heap.Push(hp, Rectangle{1, 1})
	fmt.Println((*hp)[0], hp)
	fmt.Println(hp)
	heap.Pop(hp)
	fmt.Println(hp)
}
