package model

import "time"

type DataNode struct {
	Id        int
	alive     int // 0 died ; 1 alive ; 2 waiting
	waittimer *time.Timer
	dietimer  *time.Timer
}
