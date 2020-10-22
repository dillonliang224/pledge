package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "liang", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
