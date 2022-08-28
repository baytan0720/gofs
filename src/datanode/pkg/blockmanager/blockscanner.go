package blockmanager

import (
	"errors"
	"gofs/src/service"
	"log"
	"os"
	"path/filepath"
)

func ScanBlocks(path string) []*service.BlockInfo {
	blocklist := make([]*service.BlockInfo, 0, 1000)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return errors.New("empty")
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
	return blocklist[1:]
}
