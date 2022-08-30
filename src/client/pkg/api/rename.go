package api

import (
	"context"
	"fmt"
	"gofs/src/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Rename(gofspath, filename, Addr string) {
	err := checkPath(gofspath)
	if err != nil {
		fmt.Println("Invalid Argument: ", err)
		return
	}
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("NameNode connect fail:", err)
	}
	c := service.NewNameNodeServiceClient(conn)
	rep, _ := c.Rename(context.Background(), &service.RenameArgs{Path: gofspath, NewName: filename})
	if rep.Status == service.StatusCode_NotOK {
		fmt.Print("Rename fail:")
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
	fmt.Println("Sucess")
}
