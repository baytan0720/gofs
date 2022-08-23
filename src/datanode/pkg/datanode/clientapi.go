package datanode

import (
	"context"
	"crypto/md5"
	"gofs/src/datanode/pkg/blockmanager"
	"gofs/src/service"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const hextable = "0123456789abcdef"

var pipeline map[int32]*grpc.ClientConn
var timer *time.Timer

func (dn *DataNode) WriteBlock(ctx context.Context, args *service.WriteBlockArgs) (*service.WriteBlockReply, error) {
	if fastMD5(args.DataBuf) != args.Md5 {
		return &service.WriteBlockReply{Status: service.StatusCode_NotOK}, nil
	}
	if args.Index != 2 {
		args.Index++
		go func() {
			conn := pipeline[args.DatanodeIds[args.Index]]
			c := service.NewDataNodeServiceClient(conn)
			c.WriteBlock(context.Background(), args)
			timer.Reset(0)
			conn.Close()
		}()
	}
	go func(path, blockid string, data []byte) {
		err := blockmanager.WriteBlock(path, blockid, data)
		if err != nil {
			return
		}
		dn.newBlockReport(blockid, args.Size, args.EntryId)
	}(dn.BlockPath, args.BlockId, args.DataBuf)
	return &service.WriteBlockReply{Status: service.StatusCode_OK}, nil
}

func (dn *DataNode) CreatePipeline(ctx context.Context, args *service.CreatePipelineArgs) (*service.CreatePipelineReply, error) {
	if args.Index != 2 {
		args.Index++
		conn, err := grpc.Dial(args.DataNodes[args.Index+1].Addr+args.DataNodes[args.Index+1].Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return &service.CreatePipelineReply{Status: service.StatusCode_NotOK}, nil
		}
		c := service.NewDataNodeServiceClient(conn)
		rep, err := c.CreatePipeline(context.Background(), args)
		if err != nil || rep.Status != service.StatusCode_OK {
			return &service.CreatePipelineReply{Status: service.StatusCode_NotOK}, nil
		}
		pipeline[args.DataNodes[args.Index].Id] = conn
	}
	timer = time.NewTimer(10 * time.Minute)
	go func(id int32) {
		<-timer.C
		pipeline[id].Close()
	}(args.DataNodes[args.Index].Id)
	return &service.CreatePipelineReply{Status: service.StatusCode_OK}, nil
}

func fastMD5(data []byte) string {
	src := md5.Sum(data)
	var dst = make([]byte, 32)
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return string(dst)
}
