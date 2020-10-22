package utils

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func TestString2bytes(t *testing.T) {
	var s = "liang"
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	t.Log(spew.Sdump(sh))

	b := []byte(s)
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	t.Log(spew.Sdump(bh))

	bNoAlloc := String2bytes(s)
	bhNoAlloc := (*reflect.SliceHeader)(unsafe.Pointer(&bNoAlloc))
	t.Log(spew.Sdump(bhNoAlloc))

	if sh.Data != bhNoAlloc.Data {
		t.Fail()
	}
}

func TestBytes2string(t *testing.T) {
	b := []byte("liang")
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	t.Log(spew.Sdump(bh))

	str := string(b)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	t.Log(spew.Sdump(sh))

	strNoAlloc := Bytes2string(b)
	shNoAlloc := (*reflect.StringHeader)(unsafe.Pointer(&strNoAlloc))
	t.Log(spew.Sdump(shNoAlloc))

	if bh.Data != shNoAlloc.Data {
		t.Fail()
	}
}
