package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello: " + request.GetValue()
	return nil
}

func main() {
	_ = rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
