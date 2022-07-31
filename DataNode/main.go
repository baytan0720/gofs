package main

import (
	"log"
	"net/rpc"
)

type DataNode struct {
	Id int
}

func main() {
	register()

	select {}
}

func register() {
	nn, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}

	Args := &Args{}
	Reply := &Reply{}
	err = nn.Call("NameNode.DNRegister", &Args, &Reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	log.Println("Register Success")
}
