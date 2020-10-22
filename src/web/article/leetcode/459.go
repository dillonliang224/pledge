package main

import (
	"strings"
)

func main() {
	r := repeatedSubstringPattern("abcdabcd")
	print(r)
}

func repeatedSubstringPattern(s string) bool {
	// b := []byte(s + s)
	return strings.Contains(string([]byte(s + s)[1:2*len(s)-1]), s)
}
