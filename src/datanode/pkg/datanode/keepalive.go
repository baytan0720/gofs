package datanode

import (
	"context"
	"gofs/src/service"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (dn *DataNode) HeartBeat() {
	for {
		c := service.NewNameNodeServiceClient(dn.conn)
		rep, err := c.HeartBeat(context.Background(), &service.HeartBeatArgs{Id: int32(dn.id)})
		if err != nil {
			log.Println(err)
			dn.reconnect()
			continue
		}
		if rep.Status == service.StatusCode_NotRegister {
			dn.Register()
		}
		time.Sleep(3 * time.Second)
	}
}

func (dn *DataNode) reconnect() {
	for i := 0; i < 20; i++ {
		addr := dn.NameNodeAddr + dn.NameNodePort
		conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		dn.conn.Close()
		dn.conn = conn
		c := service.NewNameNodeServiceClient(conn)
		dn.client = c
		_, err := dn.client.HeartBeat(context.Background(), &service.HeartBeatArgs{Id: int32(dn.id)})
		if err == nil {
			return
		}
		time.Sleep(3 * time.Second)
	}
	log.Fatal("Connect timeout: DataNode offline")
}
