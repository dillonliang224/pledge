package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"git.dillonliang.cn/micro-svc/pledge/src/web/article/base"
)

type Message struct {
}

// 逆序+输出
func main() {
	head := base.NewLNode()
	base.CreateNode(head, 10)
	base.PrintNode("before: ", head)

	// 就地逆序
	// Reverse(head)

	// 插入逆序
	// InsertReverse(head)

	// 递归
	// RecursiveReverse(head)

	// 逆序+顺序输出
	base.ReverseSortV2(head)

	// 递归输出
	// RecursiveReversePrint(head)

	// 就地逆序并输出，改变了原链表的结构
	base.PrintNode("after: ", head)

	m := map[string]interface{}{
		"dillon": []byte("liang"),
		"age":    1,
	}
	d, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(d))

	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n %v\n", w, w)
	fmt.Println(w == nil)
}
