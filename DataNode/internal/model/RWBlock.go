package model

import (
	"os"
	"time"
)

func (dn *DataNode) WriteBlock(block, blockid []byte) {
	time.Sleep(10 * time.Second)
	file, err := os.OpenFile("../blocks/"+string(blockid), os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	// 写入字节流
	_, err = file.Write(block)
	if err != nil {
		return
	}
	dn.Blockreport()
}

func (dn *DataNode) ReadBlock(blockid []byte) []byte {
	file, err := os.Open("../blocks/" + string(blockid))
	if err != nil {
		return nil
	}
	b := make([]byte, 4096)
	file.Read(b)
	return b
}
