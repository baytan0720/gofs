package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
}

type Reply struct {
	S string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
		return
	}

	Args := &Args{}
	Reply := &Reply{}
	err = client.Call("NameNode.Hello", &Args, &Reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(Reply.S)
}
