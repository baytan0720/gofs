package model

import (
	"context"
	"gofs/DataNode/internal/service"
	"log"
	"time"
)

// Heartbeat DataNode调用心跳检测客户端
func (dn *DataNode) Heartbeat() {
	for {
		c := service.NewHeartbeatServiceClient(dn.Conn)
		res, err := c.Heartbeat(context.Background(), &service.HeartbeatArgs{Id: int32(dn.Id)})
		if err != nil {
			log.Println("Connection interruption:", err, "Try to reconnect...")
			timer := time.NewTimer(time.Minute)
			go func() {
				<-timer.C
				log.Fatal("DataNode offline")
			}()
			dn.reconnect()
			timer.Stop()
			continue
		}
		if res.ACK == 0 {
			conn, id, addr := DNRegister()
			dn.Conn = conn
			dn.Id = id
			dn.Addr = addr
			continue
		}
		time.Sleep(3 * time.Second)
	}
}
