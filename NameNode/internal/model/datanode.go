package model

import (
	"gofs/NameNode/internal/service"
	"time"
)

type DataNode struct {
	Id        int
	alive     int // 0 died ; 1 alive ; 2 waiting
	Blocklist []*service.Block
	waittimer *time.Timer
	dietimer  *time.Timer
}
