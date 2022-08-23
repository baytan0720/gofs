package blockmanager

import (
	"errors"
	"gofs/src/service"
	"log"
	"os"
	"path/filepath"
)

func ScanBlocks(path string) ([]*service.BlockInfo, int64) {
	blocklist := make([]*service.BlockInfo, 0, 1000)
	var used int64
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return errors.New("empty")
		}
		if len(blocklist) > 0 {
			used += info.Size()
		}
		blocklist = append(blocklist, &service.BlockInfo{
			Id:   info.Name(),
			Size: info.Size(),
		})
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return blocklist[1:], used
}
