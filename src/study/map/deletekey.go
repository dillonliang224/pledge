package main

import (
	"sync"
)

func main() {
	var sm sync.Map
	var value [16]byte

	for i := 0; i < 1<<26; i++ {
		sm.Store(i, value)
		sm.Delete(i - 1)
	}
}
