package model

import (
	"context"
	"fmt"
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

	service.UnimplementedHeartbeatServiceServer
	service.UnimplementedRegisterServiceServer
	service.UnimplementedBlockreportServiceServer
}

//MakeNameNode 创建NameNode
func MakeNameNode() *NameNode {
	config.Opencfg()

	nn := NameNode{
		idChan:       make(chan int, 10),
		DataNodeList: make([]*DataNode, 20, config.Config.NumDataNodeLimit),
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
	waittimer := time.NewTimer(10 * time.Second)
	dietimer := time.NewTimer(1 * time.Minute)
	nn.DataNodeList[rep.Id] = &DataNode{
		Id:        int(rep.Id),
		alive:     1,
		waittimer: waittimer,
		dietimer:  dietimer,
	}

	dietimer.Stop()
	go func() {
		for {
			<-waittimer.C
			nn.DataNodeList[rep.Id].alive = 2
			log.Println("ID: ", rep.Id, " is waiting reconnect")
			waittimer.Stop()
			dietimer.Reset(1 * time.Minute)
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
	rep := &service.HeartbeatReply{
		ACK: 1,
	}
	if nn.DataNodeList[args.Id] == nil {
		rep.ACK = 0
		return rep, nil
	}
	dn := nn.DataNodeList[args.Id]
	dn.waittimer.Stop()
	dn.dietimer.Stop()
	dn.waittimer.Reset(10 * time.Second)
	if dn.alive == 2 {
		dn.alive = 1
	}
	// log.Println("ID: ", args.Id, " Heartbeating")
	return rep, nil
}

func (nn *NameNode) Blockreport(ctx context.Context, args *service.BlockreportArgs) (*service.BlockreportReply, error) {
	nn.DataNodeList[args.Id].Blocklist = args.Blocklist
	log.Println()
	for _, v := range args.Blocklist {
		fmt.Println("Filename:", v.Name, "Size:", v.Size, "B", "Modtime:", v.Modtime)
	}
	return &service.BlockreportReply{ACK: 1}, nil
}

func (nn *NameNode) getId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
}
