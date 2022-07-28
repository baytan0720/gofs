package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type NameNode struct{}

func main() {
	nn := makeNameNode()
	nn.server()

	//阻塞
	select {}
}

//rpc调用示例
func (nn *NameNode) Hello(Args *Args, Reply *Reply) error {
	Reply.S = "Hello"
	return nil
}

//创建NameNode
func makeNameNode() *NameNode {
	nn := NameNode{}
	return &nn
}

//rpc挂载
func (nn *NameNode) server() {
	rpc.Register(nn)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// func (nn *NameNode) close() {

// }
