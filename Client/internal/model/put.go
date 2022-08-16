package model

import (
	"context"
	"crypto/md5"
	"fmt"
	"gofs/Client/config"
	"gofs/Client/internal/service"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Put(gofspath, inputpath string) {
	//读取上传的文件
	file, err := os.Open(inputpath)
	if err != nil {
		fmt.Println("Cannot Open Inputfile:", err)
		return
	}
	defer file.Close()
	fileinfo, _ := file.Stat()

	//向NameNode请求上传文件，入参为远端路径和文件信息，返回ACK 0/1以及block的size，block的数目
	addr := config.Config.Addr + config.Config.Port
	nnrpcconn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := service.NewPutFileServiceClient(nnrpcconn)
	ctx := context.Background()
	res, err := c.PutFile(ctx, &service.PutFileArgs{Path: gofspath, FileName: file.Name(), Size: uint64(fileinfo.Size())})
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.ACK != 1 {
		fmt.Println(res.Err)
		return
	}
	blocksize := int64(res.BlockSize)
	blocksum := len(res.BlockId)
	blockId := res.BlockId

	//读取block
	b := make([]byte, blocksize)
	index := 0
	var i int64 = 1
	for ; i <= int64(blocksum); i++ {
		//设置偏移量
		file.Seek((i-1)*(int64(blocksize)), 0)
		if len(b) > int((fileinfo.Size() - (i-1)*blocksize)) {
			b = make([]byte, fileinfo.Size()-(i-1)*blocksize)
		}
		file.Read(b)
		Md5 := md5.Sum(b)

		//向NameNode请求上传block，返回ACK 0/1以及DataNode的信息
		c := service.NewPutBlockServiceClient(nnrpcconn)
		res, err := c.PutBlock(ctx, &service.PutBlockArgs{})
		if err != nil {
			fmt.Println(err)
			return
		}
		if res.ACK != 1 {
			fmt.Println(res.Err)
			return
		}

		//告知DataNode创建TCP服务器，入参为剩余DataNode信息，返回ACK 0/1
		dnrpcaddr := config.Config.Addr + ":" + strconv.Itoa(int(res.DNList[0].Id+1024))
		dnrpcconn, _ := grpc.Dial(dnrpcaddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		c2 := service.NewPipelineToClientServiceClient(dnrpcconn)
		ctx := context.Background()
		otherdn := make([]uint32, 2)
		otherdn[0] = res.DNList[1].Id
		otherdn[1] = res.DNList[2].Id
		res2, err := c2.PipelineToClient(ctx, &service.PipelineToClientArgs{OtherDNId: otherdn})
		if err != nil {
			fmt.Println(err)
			return
		}
		if res2.ACK != 1 {
			fmt.Println(res2.Err)
			return
		}
		socketaddr := config.Config.Addr + ":" + strconv.Itoa(int(res.DNList[0].Id+2048))
		dntcpconn, err := net.Dial("tcp", socketaddr)
		if err != nil {
			fmt.Println(err)
			return
		}

		//向DataNode写block
		dntcpconn.Write([]byte(blockId[index]))
		dntcpconn.Write(Md5[:])
		ACK := make([]byte, 1)
		dntcpconn.Read(ACK)
		if ACK[0] != '1' {
			return
		}

		dntcpconn.Write(b)
		dntcpconn.Read(ACK)
		if ACK[0] != '1' {
			fmt.Println("fail")
			return
		}

		//成功
		index++
	}
	fmt.Println("OK")
}
