package main

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

func (c *Cat) Quack() {
	println(c.Name + " meow")
}

func main() {
	c := &Cat{
		"grooming",
	}

	var d Duck
	d = c

	d.Quack()
	// var c Duck = &Cat{Name: "Dillon"}
	// switch c.(type) {
	// case *Cat:
	// 	cat := c.(*Cat)
	// 	cat.Quack()
	// }
}
