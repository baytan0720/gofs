package model

import (
	"context"
	"fmt"
	"gofs/Client/config"
	"gofs/Client/internal/service"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DNInfo() {
	addr := config.Config.Addr + config.Config.Port
	conn, _ := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	c := service.NewDNInfoServiceClient(conn)
	ctx := context.Background()
	res, err := c.DNInfo(ctx, &service.DNInfoArgs{})
	if err != nil {
		log.Fatal(err)
	}
	if len(res.DNList) == 0 {
		fmt.Println("No DataNode Online")
	}
	for _, v := range res.DNList {
		fmt.Println("DataNode"+":"+strconv.Itoa(int(v.Id))+": BlockList:", v.BlockList)
	}
}
