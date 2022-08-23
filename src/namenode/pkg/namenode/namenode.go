package namenode

import (
	"gofs/src/namenode/pkg/metamanager"
	"gofs/src/service"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/yitter/idgenerator-go/idgen"
	"google.golang.org/grpc"
)

type NameNode struct {
	BlockSize  int64
	ReplicaNum int

	NameNodeAddr     string
	NameNodePort     string
	MetaDataPath     string
	MetaDataBackup   string
	LogPath          string
	MetaDataRTime    int
	HeartBeatTimeout int

	status         int //0 safemode/1 active
	idChan         chan int32
	idIncrease     int32
	DataNodeList   []*DataNode
	DataNodeNum    int
	totaldiskquota int64
	useddisk       int64
	lease          int
	leasetimer     *time.Timer

	service.UnimplementedNameNodeServiceServer
}

func MakeNameNode() *NameNode {
	nn := &NameNode{
		status:       0,
		idChan:       make(chan int32, 3),
		idIncrease:   0,
		DataNodeList: make([]*DataNode, 3, 128),
		leasetimer:   time.NewTimer(10 * time.Minute),
	}
	nn.leasetimer.Stop()
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
	s.Serve(l)
}

func (nn *NameNode) plugin() {
	metamanager.Format(nn.MetaDataPath, nn.MetaDataBackup, nn.MetaDataRTime)
	idgen.SetIdGenerator(idgen.NewIdGeneratorOptions(1))
}

func (nn *NameNode) opencfg() {
	var path string
	if runtime.GOOS == "windows" {
		pwd, _ := os.Getwd()
		path = pwd + ""
	} else {
		path = "../../../config/config.toml"
	}
	_, err := toml.DecodeFile(path, nn)
	if err != nil {
		log.Fatal("Config Read Fail: ", err)
	}
	nn.BlockSize <<= 20
}
