package model

import (
	"context"
	"gofs/DataNode/config"
	"gofs/DataNode/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type DataNode struct {
	Id   uint32 // DataNode标识符
	Conn *grpc.ClientConn
}

// Heartbeat DataNode调用心跳检测客户端
func (dn *DataNode) Heartbeat() {
	for {
		if dn.Conn == nil {
			log.Println("Client Connection is null")
		}
		c := service.NewHeartbeatServiceClient(dn.Conn)
		_, err := c.Heartbeat(context.Background(), &service.HeartbeatArgs{Id: int32(dn.Id)})
		if err != nil {
			log.Println("Connection interruption:", err, "Try to reconnect")
			ok := false
			for i := 0; i < 10; i++ {
				ok = dn.reconnect()
				if ok {
					break
				}
				time.Sleep(3 * time.Second)
			}
		}
		time.Sleep(3 * time.Second)
	}
}

func (dn *DataNode) reconnect() bool {
	addr := config.Config.Addr + config.Config.Port
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Reconnect fail:", err)
		return false
	}
	dn.Conn = conn
	return true
}

func DNRegister() *DataNode {
	config.Opencfg()
	addr := config.Config.Addr + config.Config.Port
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := service.NewRegisterServiceClient(conn)
	req := &service.DNRegisterArgs{}
	res, err := c.Register(context.Background(), req)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Register Success,get ID: ", res.Id)
	return &DataNode{
		Id:   res.Id,
		Conn: conn,
	}
}
