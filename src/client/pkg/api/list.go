package api

import (
	"context"
	"fmt"
	"gofs/src/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func List(gofspath, Addr string) {
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
	rep, _ := c.List(context.Background(), &service.ListArgs{Path: gofspath})
	if rep.Status == service.StatusCode_NotOK {
		fmt.Print("List fail:")
		if rep.FileStatus == service.FileStatus_fPathNotFound {
			fmt.Println("path not found")
		}
		if rep.FileStatus == service.FileStatus_fIsFile {
			fmt.Println("exist file on path")
		}
		return
	}
	for _, v := range rep.Files {
		fmt.Print(v + "\t")
	}
	fmt.Println("")
}
