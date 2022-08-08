package model

import (
	"gofs/DataNode/config"
	"gofs/DataNode/internal/service"
	"log"
	"os"
	"path/filepath"
)

//扫描本地block，用于block上报
func Scanblock() []*service.Block {
	blocks := []*service.Block{}
	var root string
	if config.Config.GOOS == "windows" {
		pwd, _ := os.Getwd()
		root = pwd + "\\DataNode\\blocks"
	} else {
		root = "../blocks"
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		timeLayoutStr := "2006-01-02 15:04"
		blocks = append(blocks, &service.Block{
			Name:    info.Name(),
			Modtime: info.ModTime().Format(timeLayoutStr),
			Size:    uint64(info.Size()),
		})
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return blocks[1:]
}
