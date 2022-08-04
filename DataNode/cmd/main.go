package main

import (
	"gofs/DataNode/internal/model"
)

func main() {
	dn := model.MakeDataNode()
	go dn.Heartbeat()

	select {}
}
