package model

import (
	"context"
	"gofs/DataNode/config"
	"gofs/DataNode/internal/service"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataNode struct {
	Id   uint32 // DataNode标识符
	Conn *grpc.ClientConn
}

func MakeDataNode() *DataNode {
	conn, id := DNRegister()
	return &DataNode{
		Id:   id,
		Conn: conn,
	}
}

func DNRegister() (*grpc.ClientConn, uint32) {
	config.Opencfg()
	addr := config.Config.Addr + config.Config.Port
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := service.NewRegisterServiceClient(conn)
	req := &service.DNRegisterArgs{}
	res, err := c.Register(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Register Success,get ID: ", res.Id)
	return conn, res.Id
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
			log.Println("Connection interruption:", err, "Try to reconnect...")
			timer := time.NewTimer(time.Minute)
			go func() {
				<-timer.C
				log.Fatal("DataNode offline")
			}()
			dn.reconnect()
			timer.Stop()
		}
		time.Sleep(3 * time.Second)
	}
}

//需要重连的情况：NameNode挂了，重连并重新注册；DataNode或NameNode网络波动，不需要重新注册，重新发送心跳即可
func (dn *DataNode) reconnect() {
	addr := config.Config.Addr + config.Config.Port
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	dn.Conn.Close()
	dn.Conn = conn
}
