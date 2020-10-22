package main

import (
	"fmt"
)

type Person struct {
	age int
}

func (p Person) howOld() int {
	return p.age
}

func (p *Person) growUp() {
	p.age += 1
}

type Coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (g *Gopher) code() {
	fmt.Printf("I am coding %s language \n", g.language)
}

func (g *Gopher) debug() {
	fmt.Printf("I am debuging %s language \n", g.language)
}

func main() {
	dillon := Person{age: 28}
	fmt.Println(dillon.howOld())
	dillon.growUp()
	fmt.Println(dillon.howOld())

	liss := &Person{age: 18}
	fmt.Println(liss.howOld())
	liss.growUp()
	fmt.Println(liss.howOld())

	var c Coder = &Gopher{language: "go"}
	c.code()
	c.debug()

	// var d Coder = Gopher{language: "node.js"}

}
