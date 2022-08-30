package api

import (
	"context"
	"fmt"
	"gofs/src/service"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Get(gofspath, localpath, Addr string) {
	err := checkPath(gofspath)
	if err != nil {
		fmt.Println("Invalid Argument: ", err)
		return
	}
	ana := strings.Split(gofspath, "/")
	filename := ana[len(ana)-1]
	if filename == "" {
		filename = ana[len(ana)-2]
		localpath += filename
	} else {
		localpath += "/" + filename
	}

	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(66<<20), grpc.MaxCallSendMsgSize(66<<20))
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("NameNode connect fail:", err)
		return
	}
	defer conn.Close()
	c := service.NewNameNodeServiceClient(conn)

	rep, err := c.GetFile(context.Background(), &service.GetFileArgs{Path: gofspath})
	if err != nil {
		fmt.Println("Get fail:", err)
		return
	} else if rep.Status == service.StatusCode_NotOK {
		fmt.Print("Get fail: ")
		if rep.FileStatus == service.FileStatus_fPathNotFound {
			fmt.Println("path not found")
		}
		if rep.FileStatus == service.FileStatus_fIsFile {
			fmt.Println("exist file on path")
		}
		return
	}

	file, err := os.OpenFile(localpath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Can't write file on local:", err)
		return
	}
	defer file.Close()

	databuf := make([]chan []byte, len(rep.Blocks))
	for i, v := range rep.Blocks {
		databuf[i] = make(chan []byte)
		for _, j := range v.DataNodes {
			go func(blockid string, index int, info *service.DataNodeNetInfo) {
				conn, err := grpc.Dial(info.Addr+info.Port, grpc.WithTransportCredentials(insecure.NewCredentials()), diaOpt)
				if err != nil {
					fmt.Println("DataNode connect fail:", err)
					return
				}
				c := service.NewDataNodeServiceClient(conn)
				defer conn.Close()

				rep, err := c.ReadBlock(context.Background(), &service.ReadBlockArgs{BlockId: blockid})
				if err != nil {
					fmt.Println("DataNode connect fail:", err)
					return
				}
				if rep.Status == service.StatusCode_NotOK {
					return
				}
				if fastMD5(rep.DataBuf) != rep.Md5 {
					return
				}
				defer func() {
					recover()
				}()
				if len(databuf[index]) == 0 {
					databuf[index] <- rep.DataBuf
				}
			}(v.BlockId, i, j)
		}
	}
	data := make([]byte, 0, len(databuf)*(64<<20))
	for i := 0; i < len(databuf); i++ {
		data = append(data, <-databuf[i]...)
		close(databuf[i])
	}
	file.Write(data)
}
