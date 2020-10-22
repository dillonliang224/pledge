package main

import (
	"fmt"
	"strings"
)

// 无重复的最长子串
// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcabbcc"))
	fmt.Println(lengthOfLongestSubstringV2("abcabcabbcc"))
	m := make(chan int, 0)
	fmt.Println(len(m), cap(m))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var res int
	head, tail := 0, 0
	for tail < len(s)-1 {
		tail++

		if strings.Index(s[head:tail], string(s[tail])) == -1 {
			res = max1776(tail-head+1, res)
		} else {
			for strings.Index(s[head:tail], string(s[tail])) > 0 {
				head++
			}
		}
	}
	// low, fast := 0, 1
	// for i := 0; i < len(s)-1; i++ {
	// 	for j := low; j < fast; j++ {
	// 		if s[j] == s[fast] {
	// 			low = j + 1
	// 		}
	// 	}
	//
	// 	res = max1776(res, fast-low+1)
	// 	fast++
	// }

	return res
}

func lengthOfLongestSubstringV2(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	m := make(map[byte]bool)
	left := 0
	m[s[0]] = true
	res := 1
	for right := 1; right < len(s); {
		if exists := m[s[right]]; exists {
			delete(m, s[left])
			left++
		} else {
			m[s[right]] = true
			right++
		}

		if right-left > res {
			res = right - left
		}
	}

	return res
}

func max1776(a, b int) int {
	if a > b {
		return a
	}
	return b
}
