package datanode

import (
	"context"
	"gofs/src/datanode/pkg/blockmanager"
	"gofs/src/service"
	"log"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (dn *DataNode) UpdateLoad() {
	go func() {
		var temp []float64
		for {
			temp, _ = cpu.Percent(0, false)
			dn.load.Percpu = float32(temp[0])
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		var stat *mem.VirtualMemoryStat
		for {
			stat, _ = mem.VirtualMemory()
			dn.load.Permem = float32(stat.UsedPercent)
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		var stat *disk.UsageStat
		for {
			stat, _ = disk.Usage("/")
			dn.load.TotalDisk = int64(stat.Total)
			dn.load.UsedDisk = int64(stat.Used)
			dn.load.Perdisk = float32(stat.UsedPercent)
			time.Sleep(3 * time.Second)
		}
	}()
}

func (dn *DataNode) HeartBeat() {
	for {
		c := service.NewNameNodeServiceClient(dn.conn)
		rep, err := c.HeartBeat(context.Background(), &service.HeartBeatArgs{Id: dn.id, Load: dn.load})
		if err != nil {
			log.Println(err)
			dn.reconnect()
			continue
		}
		if rep.Status == service.StatusCode_NotRegister {
			dn.Register()
		}
		for _, v := range rep.CleanBlockId {
			go blockmanager.DeleteBlock(dn.BlockPath, v)
		}
		time.Sleep(3 * time.Second)
	}
}

func (dn *DataNode) reconnect() {
	for i := 0; i < 10; i++ {
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
