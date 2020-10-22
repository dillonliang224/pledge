package main

import (
	"os"
)

func main() {
	reader, writer, _ := os.Pipe()
	defer writer.Close()

	_, _ = writer.Write([]byte("hello world"))
	r := make([]byte, 100)
	_, _ = reader.Read(r)

	// _, _ = writer.Write([]byte("hello world"))

	print(string(r))

}
