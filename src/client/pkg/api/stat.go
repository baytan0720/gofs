package api

import (
	"context"
	"fmt"
	"gofs/src/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Stat(gofspath string) {
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
	rep, _ := c.Stat(context.Background(), &service.StatArgs{Path: gofspath})
	if rep.Status == service.StatusCode_NotOK {
		fmt.Print("Stat fail:")
		if rep.FileStatus == service.FileStatus_fIsFile {
			fmt.Println("exist file on path")
		}
		if rep.FileStatus == service.FileStatus_fPathNotFound {
			fmt.Println("path not found")
		}
		return
	}
	if rep.Info.FileType == service.FileStatus_fIsDirectory {
		fmt.Print("Type: Directory\t")
	} else {
		fmt.Print("Type: File\t")
		fmt.Printf("Size: %d\t", rep.Info.Size)
	}
	fmt.Printf("Modtime: %s\n", rep.Info.Modtime)
}
