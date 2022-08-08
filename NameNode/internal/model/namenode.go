package model

import (
	"gofs/NameNode/config"
	"gofs/NameNode/internal/service"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type NameNode struct {
	NumDataNode  int
	idChan       chan int
	idIncrease   int
	DataNodeList []*DataNode
	mu           *sync.Mutex

	service.UnimplementedHeartbeatServiceServer
	service.UnimplementedRegisterServiceServer
	service.UnimplementedBlockreportServiceServer
}

//MakeNameNode 创建NameNode
func MakeNameNode() *NameNode {
	config.Opencfg()

	nn := NameNode{
		idChan:       make(chan int, 10),
		DataNodeList: make([]*DataNode, 16, 128),
		mu:           &sync.Mutex{},
	}
	return &nn
}

//rpc挂载
func (nn *NameNode) Server() {
	l, e := net.Listen("tcp", config.Config.Port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	s := grpc.NewServer()
	service.RegisterRegisterServiceServer(s, nn)
	service.RegisterHeartbeatServiceServer(s, nn)
	service.RegisterBlockreportServiceServer(s, nn)
	log.Println("NameNode is running, listen on " + "127.0.0.1" + config.Config.Port)
	s.Serve(l)
}
