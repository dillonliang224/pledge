package main

import (
	"fmt"
)

func main() {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range five {
		five = five[:2]
		fmt.Println(five)
		fmt.Printf("v[%s]\n", v)
	}
}
