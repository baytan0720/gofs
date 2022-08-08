package model

import (
	"crypto/md5"
	"fmt"
	"gofs/DataNode/internal/service"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// WriteDisk 负责将文件落盘 port 是为了防止命名重复
func WriteDisk(req *service.ReplicationFile, port string) (*service.Block, error) {
	has := md5.Sum(req.Payload)
	md5str := fmt.Sprintf("%x", has)
	md5str = md5str[9:25]
	// TODO block差一个Sequence属性
	block := &service.Block{
		Name:    md5str,
		Modtime: time.Now().String(),
		Size:    req.Length,
	}
	path := fmt.Sprintf("./DataNode/blocks/%s_%s", md5str, port)
	if runtime.GOOS == "windows" {
		path = fmt.Sprintf(".\\DataNode\\blocks\\%s_%s", md5str, port)
	}
	err := os.WriteFile(path, req.Payload, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return block, nil
}

// ReadDisk 负责将读取文件数据 port 是为了防止命名重复
func ReadDisk(name string, port string) (*service.ReplicationFile, error) {
	name = strings.Join([]string{name, port}, "_")
	path := fmt.Sprintf("./DataNode/blocks/%s", name)
	if runtime.GOOS == "windows" {
		path = fmt.Sprintf(".\\DataNode\\blocks\\%s", name)
	}
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &service.ReplicationFile{
		Payload: b,
		Length:  uint64(len(b)),
	}, nil
}
