package model

import (
	"context"
	"gofs/NameNode/config"
	"gofs/NameNode/internal/service"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type NameNode struct {
	NumDataNode  int
	idChan       chan int
	idIncrease   int
	DataNodeList []*DataNode
	mu           *sync.Mutex
}

//MakeNameNode 创建NameNode
func MakeNameNode() *NameNode {
	config.Opencfg()

	nn := NameNode{
		idChan:       make(chan int, 10),
		DataNodeList: make([]*DataNode, 10, config.Config.NumDataNodeLimit),
		mu:           &sync.Mutex{},
	}
	return &nn
}

// Register 由DataNode调用该方法，将对应DataNode注册到本NameNode上
func (nn *NameNode) Register(ctx context.Context, args *service.DNRegisterArgs) (*service.DNRegisterReply, error) {
	rep := new(service.DNRegisterReply)
	select {
	case t := <-nn.idChan:
		rep.Id = uint32(t)
	default:
		nn.getId()
		rep.Id = uint32(<-nn.idChan)
	}
	nn.mu.Lock()
	nn.NumDataNode++
	nn.mu.Unlock()

	//定时器，10秒无心跳则等待重连，十分钟无心跳则判定离线
	waittimer := time.NewTimer(1 * time.Minute)
	dietimer := time.NewTimer(10 * time.Second)
	nn.DataNodeList[rep.Id] = &DataNode{
		Id:        int(rep.Id),
		alive:     1,
		waittimer: waittimer,
		dietimer:  dietimer,
	}

	go func() {
		for {
			<-waittimer.C
			nn.DataNodeList[rep.Id].alive = 2
			log.Println("ID: ", rep.Id, " is waiting reconnect")
			waittimer.Stop()
		}
	}()
	go func() {
		<-dietimer.C
		nn.DataNodeList[rep.Id] = nil
		dietimer.Stop()
		nn.idChan <- int(rep.Id)
		log.Println("ID: ", rep.Id, " is died")
	}()

	log.Println("ID: ", rep.Id, " is connected")
	return rep, nil
}

func (nn *NameNode) Heartbeat(ctx context.Context, args *service.HeartbeatArgs) (*service.HeartbeatReply, error) {
	rep := new(service.HeartbeatReply)
	nn.DataNodeList[args.Id].dietimer.Reset(1 * time.Minute)
	nn.DataNodeList[args.Id].waittimer.Reset(10 * time.Second)
	nn.DataNodeList[args.Id].alive = 1
	log.Println("ID: ", args.Id, " Heartbeating")
	return rep, nil
}

func (nn *NameNode) getId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
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
	log.Println("NameNode is running, listen on " + "127.0.0.1" + config.Config.Port)
	s.Serve(l)
}
