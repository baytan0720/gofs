package main

import (
	"log"
	"net/rpc"
	"time"
)

type DataNode struct {
	Id int
}

var nn *rpc.Client

func main() {
	var dn *DataNode
	nn, dn = register()
	go dn.Heartbeat(HeartbeatArgs{Id: dn.Id}, HeartbeatReply{})

	select {}
}

func register() (*rpc.Client, *DataNode) {
	opencfg()

	nn, err := rpc.DialHTTP("tcp", Config.Addr+Config.Port)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	Args := &DNRegisterArgs{}
	Reply := &DNRegisterReply{}
	err = nn.Call("NameNode.DNRegister", &Args, &Reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	log.Println("Register Success,get ID: ", Reply.Id)
	return nn, &DataNode{Id: Reply.Id}
}

func (dn *DataNode) Heartbeat(Args HeartbeatArgs, Reply HeartbeatReply) {
	for {
		err := nn.Call("NameNode.Heartbeat", &Args, &Reply)
		if err != nil {
			log.Println("Connection interruption:", err, "Try to reconnect")
			ok := false
			for i := 0; i < 10; i++ {
				ok = reconnect()
				if ok {
					break
				}
				time.Sleep(3 * time.Second)
			}
		}

		time.Sleep(3 * time.Second)
	}
}

func reconnect() bool {
	rec, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Println("Reconnect fail:", err)
		return false
	}
	nn = rec
	return true
}
