package main

func main() {
	// 二叉搜索树的后序遍历
	postorder := []int{1, 3, 2, 6, 5}
	r := juage(postorder, 0, len(postorder)-1)
	print(r)
}

func juage(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	var i int
	for i = start; i < end; i++ {
		if postorder[i] > postorder[end] {
			break
		}
	}

	for j := i; j < end; j++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}

	return juage(postorder, start, i-1) && juage(postorder, i, end-1)
}
