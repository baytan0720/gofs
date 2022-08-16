package model

import (
	"context"
	"fmt"
	"gofs/NameNode/internal/service"
	"log"
	"time"
)

type DataNode struct {
	Id        int
	alive     int // 0 died ; 1 alive ; 2 waiting
	Blocklist []*service.Block
	waittimer *time.Timer
	dietimer  *time.Timer
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

//接收来自DataNode的心跳，重置超时定时器
func (nn *NameNode) HeartBeat(ctx context.Context, args *service.HeartBeatArgs) (*service.HeartBeatReply, error) {
	rep := &service.HeartBeatReply{
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

//接收来自DataNode的block上报
func (nn *NameNode) BlockReport(ctx context.Context, args *service.BlockReportArgs) (*service.BlockReportReply, error) {
	nn.DataNodeList[args.Id].Blocklist = args.Blocklist
	log.Println()
	for _, v := range args.Blocklist {
		fmt.Println("Filename:", v.Id, "Size:", v.Size, "B", "Modtime:", v.Modtime)
	}
	return &service.BlockReportReply{ACK: 1}, nil
}

//向idchan放入id
func (nn *NameNode) getId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
	if nn.idIncrease > len(nn.DataNodeList)-1 {
		nn.DataNodeList = append(nn.DataNodeList, make([]*DataNode, len(nn.DataNodeList))...)
	}
}
