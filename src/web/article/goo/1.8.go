package main

import "git.dillonliang.cn/micro-svc/pledge/src/web/article/base"

func main() {
	head := base.NewLNode()
	base.CreateNodeWithValue(head, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	base.PrintNode("before: ", head)
	// todo
	base.FanZhuanK(head, 3)
}
