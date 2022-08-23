package datanode

import (
	"context"
	"gofs/src/datanode/pkg/blockmanager"
	"gofs/src/service"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataNode struct {
	BlockSize  int64
	ReplicaNum int

	NameNodeAddr    string
	NameNodePort    string
	BlockPath       string
	BlockReportTime int

	addr      string
	port      string
	id        int32
	conn      *grpc.ClientConn
	client    service.NameNodeServiceClient
	blocklist []*service.BlockInfo
	DiskQuota int64
	UsedDisk  int64

	service.UnimplementedDataNodeServiceServer
}

func MakeDataNode() *DataNode {
	dn := &DataNode{
		addr: "127.0.0.1",
	}
	dn.opencfg()
	dn.Register()
	dn.checkblockpath()
	return dn
}

func (dn *DataNode) Register() {
	blockmanager.ScanBlocks(dn.BlockPath)
	addr := dn.NameNodeAddr + dn.NameNodePort
	dn.conn, _ = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dn.client = service.NewNameNodeServiceClient(dn.conn)
	rep, err := dn.client.Register(context.Background(), &service.RegisterArgs{
		Info: &service.DataNodeInfo{
			Addr:      dn.addr,
			StartTime: time.Now().Format("2006-01-02 15:04:05"),
			Status:    service.DataNodeStatus_dActive,
			DiskQuota: dn.DiskQuota,
			UsedDisk:  dn.UsedDisk,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	dn.id = rep.Id
	dn.addr = "127.0.0.1"
	dn.port = ":" + strconv.Itoa(int(dn.id)+1024)
	dn.BlockPath += "/" + strconv.Itoa(int(dn.id)) + "_blocks"
}

func (dn *DataNode) Server() {
	go dn.HeartBeat()
	go dn.BlockReport()

	l, err := net.Listen("tcp", dn.port)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	maxSize := 136314880
	s := grpc.NewServer(grpc.MaxRecvMsgSize(maxSize), grpc.MaxSendMsgSize(maxSize))
	service.RegisterDataNodeServiceServer(s, dn)
	log.Println("DataNode is running, listen on", dn.addr+dn.port)
	s.Serve(l)
}

func (dn *DataNode) opencfg() {
	var path string
	if runtime.GOOS == "windows" {
		pwd, _ := os.Getwd()
		path = pwd + ""
	} else {
		path = "../../../config/config.toml"
	}
	_, err := toml.DecodeFile(path, dn)
	if err != nil {
		log.Fatal("Config Read Fail: ", err)
	}
	dn.BlockSize <<= 20
	dn.DiskQuota <<= 30
}
