package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println(strings.ContainsRune("", 97))
	fmt.Println(strings.Split("10010111001", "0"))
	temp := strings.Split("10010111001", "0")
	fmt.Println(len(temp), temp)
	for i, v := range temp {
		fmt.Println(i, v)
	}
	fmt.Println(strings.Fields("   foo bar    bza "))
}
