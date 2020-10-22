package main

import (
	"fmt"
	"math"
)

// 如何求解最小三元组距离
func main() {
	a := []int{3, 4, 5, 7, 15}
	b := []int{10, 12, 14, 16, 17}
	c := []int{20, 21, 23, 24, 37, 30}
	find48ByManli(a, b, c)
	find48ByLen(a, b, c)
}

func find48ByLen(a, b, c []int) {
	aLen := len(a)
	bLen := len(b)
	cLen := len(c)

	var j int = 0
	var k int = 0
	var l int = 0

	min := math.MaxInt64

	for true {
		curDist := getMax3(abs(a[j]-b[k]), abs(a[j]-c[l]), abs(b[k]-c[l]))
		if curDist < min {
			min = curDist
		}

		minEle := getMin3(a[j], b[k], c[l])
		if minEle == a[j] {
			j++
			if j >= aLen {
				break
			}
		} else if minEle == b[k] {
			k++
			if k >= bLen {
				break
			}
		} else {
			l++
			if l >= cLen {
				break
			}
		}
	}

	fmt.Println("最短距离法： ", min)
}

// 蛮力求解
func find48ByManli(a, b, c []int) {
	min := getMax3(abs(a[0]-b[0]), abs(a[0]-c[0]), abs(b[0]-c[0]))
	dist := 0
	for _, ad := range a {
		for _, bd := range b {
			for _, cd := range c {
				dist = getMax3(abs(ad-bd), abs(bd-cd), abs(ad-cd))
				if dist < min {
					min = dist
				}
			}
		}
	}

	fmt.Println("暴力求解： ", min)
}

func getMax3(t1, t2, t3 int) int {
	temp := -1
	if t1 < t2 {
		temp = t2
	} else {
		temp = t1
	}

	if temp < t3 {
		return t3
	} else {
		return temp
	}
}

func getMin3(t1, t2, t3 int) int {
	temp := -1
	if t1 > t2 {
		temp = t2
	} else {
		temp = t1
	}

	if temp > t3 {
		return t3
	} else {
		return temp
	}
}

func abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0
	}
	return x
}
