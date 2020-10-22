package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 10) + " dillon"
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	b := Binary(200)
	any := fmt.Stringer(b)
	val := reflect.ValueOf(&any)
	// val.Elem().SetUint()
	fmt.Println(any, val.Elem())
}
