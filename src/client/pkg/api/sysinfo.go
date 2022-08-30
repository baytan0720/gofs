package api

import (
	"context"
	"errors"
	"fmt"
	"gofs/src/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SysInfo(Addr string) {
	conn, err := grpc.Dial(Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("NameNode connect fail:", err)
	}
	c := service.NewNameNodeServiceClient(conn)
	rep, _ := c.GetSystemInfo(context.Background(), &service.GetSystemInfoArgs{})
	fmt.Println("System Info:")
	fmt.Println("  Status:", rep.NnStatus)
	fmt.Println("  BlockNum:", rep.ReplicaNum)
	fmt.Println("  BlockSize:", rep.BlockSize>>20, "Mb")
	fmt.Println("DataNode:")
	fmt.Println("  Id\tAddr\t\tPort\tStartTime\t\tStatus\t\tTotalDisk\tUsedDisk\tBlockNum\tBlocksId\t\tBlocksSize")
	for _, v := range rep.DataNodes {
		fmt.Printf("  %d\t%s\t%s\t%s\t%v\t\t%d Gb\t\t%d %%\t\t%d\t\t", v.Id, v.Addr, v.Port[1:], v.StartTime, v.Status, v.TotalDisk>>30, v.UsedDisk/v.TotalDisk, len(v.Blocks))
		if len(v.Blocks) == 0 {
			fmt.Println()
		}
		for i, v := range v.Blocks {
			if i >= 1 {
				for i := 0; i < 15; i++ {
					fmt.Print("\t")
				}
			}
			fmt.Printf("%s\t%d B\n ", v.Id, v.Size)
		}
	}
}

func checkPath(path string) error {
	if []byte(path)[0] != '/' {
		return errors.New("please take the root '/' as the starting gofspath")
	}
	return nil
}
