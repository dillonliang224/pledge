package main

import (
	"fmt"
)

var (
	num int
	vis []bool
)

// 锁和房子
func main() {
	rooms := [][]int{
		{1},
		{2},
		{0},
	}
	// r := canVisitAllRooms(rooms)
	// fmt.Println(r)

	r2 := bfs(rooms)
	fmt.Println(r2)
}

func bfs(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)

	queue := []int{}
	vis[0] = true
	queue = append(queue, 0)
	for i := 0; i < len(queue); i++ {
		x := queue[i]
		fmt.Println(x)
		num++
		for _, t := range rooms[x] {
			if !vis[t] {
				vis[t] = true
				queue = append(queue, t)
			}
		}
	}

	return num == n
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	vis[x] = true
	num++
	for _, t := range rooms[x] {
		if !vis[t] {
			dfs(rooms, t)
		}
	}
}
