package namenode

import (
	"gofs/src/namenode/pkg/leasemanager"
	"gofs/src/namenode/pkg/logmanager"
	"gofs/src/namenode/pkg/metadatamanager"
	"gofs/src/service"
	"log"
	"net"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/yitter/idgenerator-go/idgen"
	"google.golang.org/grpc"
)

type NameNode struct {
	BlockSize  int64
	ReplicaNum int

	NameNodeAddr     string
	NameNodePort     string
	MetaDataPath     string
	LogPath          string
	HeartBeatTimeout int
	MaxLoad          int

	Status       int //0 safemode/1 active
	Starttime    string
	idChan       chan int32
	idIncrease   int32
	DataNodeList []*datanode
	DataNodeNum  int
	cache        map[string][]int
	lease        *leasemanager.Lease

	service.UnimplementedNameNodeServiceServer
}

func MakeNameNode() *NameNode {
	nn := &NameNode{
		Status:       0,
		idChan:       make(chan int32, 64),
		idIncrease:   0,
		DataNodeList: make([]*datanode, 3, 128),
		cache:        make(map[string][]int),
		Starttime:    time.Now().Format("2006-01-02 15:04:05"),
	}
	nn.opencfg()
	nn.plugin()
	return nn
}

func (nn *NameNode) Server() {
	l, err := net.Listen("tcp", nn.NameNodePort)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	s := grpc.NewServer()
	service.RegisterNameNodeServiceServer(s, nn)
	log.Println("NameNode is running, listen on " + "127.0.0.1" + nn.NameNodePort)
	logrus.Info("NameNode is running")
	s.Serve(l)
}

func (nn *NameNode) opencfg() {
	path := "../../../config/config.toml"
	_, err := toml.DecodeFile(path, nn)
	if err != nil {
		log.Fatal("Config Read Fail: ", err)
	}
	nn.BlockSize <<= 20
}

func (nn *NameNode) plugin() {
	logmanager.Start(nn.LogPath, nn.Starttime)
	metadatamanager.Start(nn.MetaDataPath)
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(1))
	nn.lease = leasemanager.MakeLease()
	nn.monitorServer()
}
