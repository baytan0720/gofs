package main

import (
	"gofs/DataNode/internal/model"
)

func main() {
	dn := model.DNRegister()
	go dn.Heartbeat()
	select {}
}
