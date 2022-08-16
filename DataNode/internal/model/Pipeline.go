package model

import (
	"context"
	"crypto/md5"
	"gofs/DataNode/internal/service"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (dn *DataNode) PipelineToClient(ctx context.Context, args *service.PipelineToClientArgs) (*service.PipelineToClientReply, error) {
	s, err := net.Listen("tcp", dn.tcpAddr)
	if err != nil {
		return &service.PipelineToClientReply{ACK: 0}, nil
	}

	otherdnconn := make([]net.Conn, 2)
	for i := 0; i < 2; i++ {
		grpcconn, _ := grpc.Dial(":"+strconv.Itoa(int(args.OtherDNId[i]+1024)), grpc.WithTransportCredentials(insecure.NewCredentials()))
		c := service.NewPipelineToDNServiceClient(grpcconn)
		c.PipelineToDN(context.Background(), &service.PipelineToDNArgs{Id: dn.Id})
		conn, err := s.Accept()
		if err != nil {
			log.Println(err)
			return &service.PipelineToClientReply{ACK: 0}, nil
		}
		log.Println(i+1, "connect")
		otherdnconn[i] = conn
		grpcconn.Close()
	}

	go func() {
		defer s.Close()
		conn, _ := s.Accept()
		defer conn.Close()
		blockId := make([]byte, 17)
		_, err := conn.Read(blockId)
		if err != nil {
			log.Println(err)
			return
		}

		correctMd5 := make([]byte, 16)
		_, err = conn.Read(correctMd5)
		if err != nil {
			log.Println(err)
			return
		}
		conn.Write([]byte{'1'})

		b := make([]byte, 4096)
		n, err := conn.Read(b)
		if err != nil {
			log.Println(err)
			return
		}
		//数据校验
		Md5 := md5.Sum(b[:n])
		for i := 0; i < 16; i++ {
			if correctMd5[i] != Md5[i] {
				conn.Write([]byte{'0'})
				return
			}
		}
		go dn.WriteBlock(b, blockId)

		for i := 0; i < 2; i++ {
			otherdnconn[i].Write(blockId)
			otherdnconn[i].Write(correctMd5)
			otherdnconn[i].Write(b[:n])
		}

		conn.Write([]byte{'1'})
	}()
	return &service.PipelineToClientReply{ACK: 1}, nil
}

func (dn *DataNode) PipelineToDN(ctx context.Context, args *service.PipelineToDNArgs) (*service.PipelineToDNReply, error) {
	conn, err := net.Dial("tcp", ":"+strconv.Itoa(int(args.Id+2048)))
	if err != nil {
		log.Println(err)
		return &service.PipelineToDNReply{ACK: 0}, nil
	}

	go func() {
		blockId := make([]byte, 17)
		_, err = conn.Read(blockId)
		if err != nil {
			log.Println(err)
			return
		}
		correctMd5 := make([]byte, 16)
		_, err = conn.Read(correctMd5)
		if err != nil {
			log.Println(err)
			return
		}

		b := make([]byte, 4096)
		n, err := conn.Read(b)
		if err != nil {
			log.Println(err)
			return
		}

		//数据校验
		Md5 := md5.Sum(b[:n])
		for i := 0; i < 16; i++ {
			if correctMd5[i] != Md5[i] {
				conn.Write([]byte{'0'})
			}
		}
		conn.Close()
	}()

	return &service.PipelineToDNReply{ACK: 1}, nil
}
