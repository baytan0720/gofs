package api

import (
	"context"
	"crypto/md5"
	"fmt"
	"gofs/src/service"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const hextable = "0123456789abcdef"

func Put(gofspath, localpath string) {
	err := checkPath(gofspath)
	if err != nil {
		fmt.Println("Invalid Argument: ", err)
		return
	}

	f, err := os.Open(localpath)
	if err != nil {
		fmt.Println("Open local file fail:", err)
		return
	}
	defer f.Close()
	finfo, _ := f.Stat()

	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("NameNode connect fail:", err)
		return
	}
	defer conn.Close()
	c := service.NewNameNodeServiceClient(conn)

	if rep, err := c.GetLease(context.Background(), &service.GetLeaseArgs{}); err != nil {
		fmt.Println("Put fail:", err)
		return
	} else if rep.Status == service.StatusCode_NotOK {
		fmt.Print("Put fail: ")
		if rep.NameNodeStatus == service.NameNodeStatus_nSafeMode {
			fmt.Println("NameNode is in safe mode")
		}
		if rep.NameNodeStatus == service.NameNodeStatus_nWritting {
			fmt.Println("Get lease refuse")
		}
		return
	} else {
		go func() {
			time.Sleep(5 * time.Minute)
			c.RenewLease(context.Background(), &service.RenewLeaseArgs{})
		}()
		defer c.ReleaseLease(context.Background(), &service.ReleaseLeaseArgs{})
	}

	rep, err := c.PutFile(context.Background(), &service.PutFileArgs{Path: gofspath, FileName: finfo.Name(), FileSize: finfo.Size()})
	if err != nil {
		fmt.Println("Put fail:", err)
		return
	}
	if rep.Status == service.StatusCode_NotOK {
		fmt.Print("Put fail: ")
		if rep.FileStatus == service.FileStatus_fExist {
			fmt.Println("file exist")
		}
		if rep.FileStatus == service.FileStatus_fIsFile {
			fmt.Println("exist file on path")
		}
		if rep.FileStatus == service.FileStatus_fPathNotFound {
			fmt.Println("path not found")
		}
		return
	}

	entryid := rep.EntryId
	blockSize := rep.BlockSize
	blockid := rep.BlockId
	var blockNum int64
	if finfo.Size()%blockSize == 0 {
		blockNum = finfo.Size() / blockSize
	} else {
		blockNum = finfo.Size()/blockSize + 1
	}
	fsize := finfo.Size()

	b := make([]byte, blockSize)
	nowBlock := 0
	var i int64 = 1
	size := blockSize
	for ; i <= int64(blockNum); i++ {
		//设置偏移量
		f.Seek((i-1)*(blockSize), 0)
		if len(b) > int((fsize - (i-1)*blockSize)) {
			b = make([]byte, fsize-(i-1)*blockSize)
			size = fsize - (i-1)*blockSize
		}
		f.Read(b)
		Md5 := fastMD5(b)

		rep, err := c.PutBlock(context.Background(), &service.PutBlockArgs{BlockSize: size})
		if err != nil {
			fmt.Println("Put fail:", err)
			return
		}
		if rep.Status == service.StatusCode_NotOK {
			fmt.Println("Put fail: Maybe have not enough DataNode")
			return
		}

		conn, err := grpc.Dial(rep.DataNodes[0].Addr+rep.DataNodes[0].Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("DataNode connect fail:", err)
			return
		}
		c := service.NewDataNodeServiceClient(conn)

		if res, err := c.CreatePipeline(context.Background(), &service.CreatePipelineArgs{Index: 0, DataNodes: rep.DataNodes}); err != nil {
			fmt.Println("Put fail:", err)
			return
		} else if res.Status != service.StatusCode_OK {
			fmt.Println("Create pipeline fail")
			return
		}

		if res, err := c.WriteBlock(context.Background(), &service.WriteBlockArgs{BlockId: blockid[nowBlock], DataBuf: b, Md5: Md5, Size: size, Index: 0, EntryId: entryid, DatanodeIds: []int32{rep.DataNodes[0].Id, rep.DataNodes[1].Id, rep.DataNodes[2].Id}}); err != nil {
			fmt.Println("DataNode connect fail:", err)
			return
		} else if res.Status != service.StatusCode_OK {
			fmt.Println("WriteBlock fail: Uknown error")
			return
		}

		conn.Close()
	}
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
