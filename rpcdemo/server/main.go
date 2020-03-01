package main

import (
	"demoCrawler/rpcdemo"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listerner, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listerner.Accept()
		if err != nil {
			log.Printf("accept error : %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)

	}
}
