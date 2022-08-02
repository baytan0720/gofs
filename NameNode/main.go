package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
)

type NameNode struct {
	NumDataNode  int
	idChan       chan int
	idIncrease   int
	DataNodeList []*DataNode
	mu           *sync.Mutex
}

func main() {
	nn := makeNameNode()
	nn.server()
	log.Println("NameNode is running, port", Config.Port)

	//阻塞
	select {}
}

//rpc调用示例
func (nn *NameNode) Hello(Args *HelloArgs, Reply *HelloReply) error {
	Reply.S = "Hello"
	return nil
}

//创建NameNode
func makeNameNode() *NameNode {
	opencfg()

	nn := NameNode{
		idChan:       make(chan int, 10),
		DataNodeList: make([]*DataNode, 10, Config.NumDataNodeLimit),
		mu:           &sync.Mutex{},
	}
	return &nn
}

//rpc挂载
func (nn *NameNode) server() {
	rpc.Register(nn)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", Config.Port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// func (nn *NameNode) close() {

// }
