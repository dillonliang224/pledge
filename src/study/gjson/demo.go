package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string `json:"name"`
	Time int64
}

func main() {
	m := Message{"dillon", 11111}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	var mm Message
	_ = json.Unmarshal(b, &mm)
	fmt.Println(mm)
}
