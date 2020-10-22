package main

import (
	"fmt"
)

func main() {
	r := isValid("{{[[()]]}}")
	fmt.Println(r)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
