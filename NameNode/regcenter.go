package main

import (
	"log"
	"time"
)

type DataNode struct {
	Id        int
	alive     int // 0 died ; 1 alive ; 2 waiting
	waittimer *time.Timer
	dietimer  *time.Timer
}

func (nn *NameNode) DNRegister(Args *DNRegisterArgs, Reply *DNRegisterReply) error {
	//请求id
	select {
	case Reply.Id = <-nn.idChan:
	default:
		nn.getId()
		Reply.Id = <-nn.idChan
	}

	nn.mu.Lock()
	nn.NumDataNode++
	nn.mu.Unlock()

	//定时器，10秒无心跳则等待重连，十分钟无心跳则判定离线
	waittimer := time.NewTimer(1 * time.Minute)
	dietimer := time.NewTimer(10 * time.Second)
	nn.DataNodeList[Reply.Id] = &DataNode{
		Id:        Reply.Id,
		alive:     1,
		waittimer: waittimer,
		dietimer:  dietimer,
	}

	go func() {
		for {
			<-waittimer.C
			nn.DataNodeList[Reply.Id].alive = 2
			log.Println("ID: ", Reply.Id, " is waiting reconnect")
			waittimer.Stop()
		}
	}()
	go func() {
		<-dietimer.C
		nn.DataNodeList[Reply.Id] = nil
		dietimer.Stop()
		nn.idChan <- Reply.Id
		log.Println("ID: ", Reply.Id, " is died")
	}()

	log.Println("ID: ", Reply.Id, " is connected")
	return nil
}

//心跳
func (nn *NameNode) Heartbeat(Args *HeartbeatArgs, Reply *HeartbeatReply) error {
	nn.DataNodeList[Args.Id].dietimer.Reset(1 * time.Minute)
	nn.DataNodeList[Args.Id].waittimer.Reset(10 * time.Second)
	nn.DataNodeList[Args.Id].alive = 1
	log.Println("ID: ", Args.Id, " Heartbeating")
	return nil
}

func (nn *NameNode) getId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
}
