package model

import (
	"context"
	"gofs/DataNode/internal/service"
	"time"
)

//block上报，告诉NameNode有哪些本地block，默认每小时一次
func (dn *DataNode) Blockreport() {
	for {
		c := service.NewBlockreportServiceClient(dn.Conn)
		_, err := c.Blockreport(context.Background(), &service.BlockreportArgs{Id: int32(dn.Id), Blocklist: dn.Blocklist})
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		time.Sleep(1 * time.Hour)
	}
}
