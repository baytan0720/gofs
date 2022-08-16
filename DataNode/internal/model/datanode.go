package model

import (
	"context"
	"gofs/DataNode/config"
	"gofs/DataNode/internal/service"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DataNode struct {
	Id        uint32 // DataNode标识符
	Blocklist []*service.Block
	Conn      *grpc.ClientConn
	grpcAddr  string //地址
	tcpAddr   string

	service.UnimplementedPipelineToClientServiceServer
	service.UnimplementedPipelineToDNServiceServer
}

func MakeDataNode() *DataNode {
	conn, id := DNRegister()
	return &DataNode{
		Id:        id,
		Blocklist: Scanblock(),
		Conn:      conn,
		grpcAddr:  ":" + strconv.Itoa(int(id+1024)),
		tcpAddr:   ":" + strconv.Itoa(int(id+2048)),
	}
}

//向NameNode注册DataNode，取得ID
func DNRegister() (*grpc.ClientConn, uint32) {
	config.Opencfg()
	addr := config.Config.Addr + config.Config.Port
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := service.NewRegisterServiceClient(conn)
	ctx := context.Background()
	res, err := c.Register(ctx, &service.DNRegisterArgs{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Register Success,get ID: ", res.Id)
	return conn, res.Id
}

//重连：NameNode挂了，重连并重新注册；DataNode或NameNode网络波动，不需要重新注册，重连并继续发送心跳即可
func (dn *DataNode) reconnect() {
	addr := config.Config.Addr + config.Config.Port
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	dn.Conn.Close()
	dn.Conn = conn
}

func (dn *DataNode) Server() {
	l, e := net.Listen("tcp", dn.grpcAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	s := grpc.NewServer()
	service.RegisterPipelineToClientServiceServer(s, dn)
	service.RegisterPipelineToDNServiceServer(s, dn)
	go dn.Heartbeat()
	go dn.Blockreport()
	s.Serve(l)
}
