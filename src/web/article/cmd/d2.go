package main

type User struct {
	Name     string
	Password string
	Age      int
}

func getUser() *User {
	a := User{}
	return &a
}

func Call1(u *User) int {
	// fmt.Printf("%v", u)
	u.Name = "dillon"
	return u.Age * 20
}

func main() {
	// a := "aaa"
	// u := &User{a, "123", 12}
	// Call1(u)

	a := make([]*int, 1)
	b := 12
	a[0] = &b
}
