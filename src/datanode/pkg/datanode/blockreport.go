package datanode

import (
	"context"
	"gofs/src/datanode/pkg/blockmanager"
	"gofs/src/service"
	"log"
	"os"
)

func (dn *DataNode) BlockReport() {
	dn.blocklist, dn.UsedDisk = blockmanager.ScanBlocks(dn.BlockPath)
	rep, err := dn.client.BlockReport(context.Background(), &service.BlockReportArgs{Blocks: dn.blocklist})
	if err != nil {
		log.Println(err)
	}
	if rep.Status == service.StatusCode_NotOK {
		log.Println("BlockReport refuse")
	}
}

func (dn *DataNode) newBlockReport(blockid string, size, entryid int64) {
	dn.blocklist = append(dn.blocklist, &service.BlockInfo{Id: blockid, Size: size})
	dn.UsedDisk += size
	if rep, err := dn.client.NewBlockReport(context.Background(), &service.NewBlockReportArgs{Id: dn.id, Block: &service.BlockInfo{Id: blockid, Size: size}, EntryId: entryid}); err != nil {
		log.Println(err)
	} else if rep.Status == service.StatusCode_NotOK {
		log.Println("NewBlockReport refuse")
	}
}

func (dn *DataNode) checkblockpath() {
	_, err := os.Stat(dn.BlockPath)
	if os.IsNotExist(err) {
		log.Println("BlockPath not found, try to mkdir")
		err := os.Mkdir(dn.BlockPath, os.ModePerm)
		if err != nil {
			log.Panic("Mkdir fail")
		}
	}
}
