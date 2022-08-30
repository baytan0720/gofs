package namenode

import (
	"context"
	"gofs/src/namenode/pkg/metadatamanager"
	"gofs/src/service"

	log "github.com/sirupsen/logrus"
)

type tempfile struct {
	filename  string
	parentid  int64
	size      int64
	blockid   []string
	cofirm    []int
	cofirmnum int
}

var WaitCofirm map[int64]*tempfile = make(map[int64]*tempfile)

func (nn *NameNode) GetSystemInfo(ctx context.Context, args *service.GetSystemInfoArgs) (*service.GetSystemInfoReply, error) {
	rep := &service.GetSystemInfoReply{
		NnStatus:   service.NameNodeStatus(nn.Status),
		ReplicaNum: int32(nn.ReplicaNum),
		BlockSize:  nn.BlockSize,
		DataNodes:  make([]*service.DataNodeInfo, 0, nn.DataNodeNum),
	}
	for _, v := range nn.DataNodeList {
		if v == nil {
			continue
		}
		v.info.TotalDisk = v.load.TotalDisk
		v.info.UsedDisk = v.load.UsedDisk
		rep.DataNodes = append(rep.DataNodes, v.info)
	}
	// metadatamanager.GetAll()
	return rep, nil
}

func (nn *NameNode) PutFile(ctx context.Context, args *service.PutFileArgs) (*service.PutFileReply, error) {
	rep := &service.PutFileReply{}
	//检查目录树
	check, parentid := metadatamanager.PutCheckPath(args.Path, args.FileName)
	if check != service.FileStatus_fPathFound {
		rep.Status = service.StatusCode_NotOK
		rep.FileStatus = check
		log.WithField("o", "Put").Error("refuse put", args.FileName, "on path", args.Path)
		return rep, nil
	}
	log.WithField("o", "Put").Info("request put", args.FileName, "on path", args.Path)
	//分配BlockId
	var blockNum int64
	if args.FileSize%nn.BlockSize == 0 {
		blockNum = args.FileSize / nn.BlockSize
	} else {
		blockNum = args.GetFileSize()/nn.BlockSize + 1
	}
	blockid := getBlockId(int(blockNum))
	rep.BlockId = blockid
	//预分配entry
	entryid := metadatamanager.NewEntryId()
	WaitCofirm[entryid] = &tempfile{blockid: blockid, cofirm: make([]int, blockNum), parentid: parentid, size: args.FileSize, filename: args.FileName}

	rep.BlockSize = nn.BlockSize
	rep.EntryId = entryid

	return rep, nil
}

func (nn *NameNode) PutBlock(ctx context.Context, args *service.PutBlockArgs) (*service.PutBlockReply, error) {
	datanodes := nn.calLoad()
	if datanodes == nil {
		log.WithField("o", "Put").Error("refule put block", args.BlockId, "to DataNodes: not enough available DataNode")
		return &service.PutBlockReply{Status: service.StatusCode_NotOK}, nil
	}
	log.WithField("o", "Put").Info("request put block", args.BlockId, "to DataNodes")
	datanodeinfo := make([]*service.DataNodeNetInfo, 3)
	for i, v := range datanodes {
		datanodeinfo[i] = &service.DataNodeNetInfo{Id: int32(v), Addr: nn.DataNodeList[v].info.Addr, Port: nn.DataNodeList[v].info.Port}
	}
	return &service.PutBlockReply{Status: service.StatusCode_OK, DataNodes: datanodeinfo}, nil
}

func (nn *NameNode) GetFile(ctx context.Context, args *service.GetFileArgs) (*service.GetFileReply, error) {
	ok, blocks := metadatamanager.Get(args.Path)
	if ok != 0 {
		log.WithField("o", "Get").Error("refuse get", args.Path, ok)
		return &service.GetFileReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "Get").Info("request get", args.Path)
	rep := &service.GetFileReply{Blocks: make([]*service.GetFileStruct, len(blocks))}
	for i, v := range blocks {
		block := &service.GetFileStruct{
			BlockId:   v,
			DataNodes: make([]*service.DataNodeNetInfo, 0, 3),
		}
		for _, c := range nn.cache[v] {
			if nn.DataNodeList[c] == nil {
				continue
			} else {
				info := nn.DataNodeList[c].info
				block.DataNodes = append(block.DataNodes, &service.DataNodeNetInfo{
					Id:   info.Id,
					Addr: info.Addr,
					Port: info.Port,
				})
			}
		}
		rep.Blocks[i] = block
	}
	return rep, nil
}

func (nn *NameNode) Mkdir(ctx context.Context, args *service.MkdirArgs) (*service.MkdirReply, error) {
	ok := metadatamanager.Mkdir(args.Path, args.DirName)
	if ok != 0 {
		log.WithField("o", "Mkdir").Error("refuse mkdir ", args.DirName, " on path", args.Path, ok)
		return &service.MkdirReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "Mkdir").Info("request mkdir ", args.DirName, " on path", args.Path)
	return &service.MkdirReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) List(ctx context.Context, args *service.ListArgs) (*service.ListReply, error) {
	files, ok := metadatamanager.List(args.Path)
	if ok != 0 {
		log.WithField("o", "List").Error("refuse list ", args.Path, ok)
		return &service.ListReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "List").Info("request list ", args.Path)
	return &service.ListReply{Files: files, Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Rename(ctx context.Context, args *service.RenameArgs) (*service.RenameReply, error) {
	ok := metadatamanager.Rename(args.Path, args.NewName)
	if ok != 0 {
		log.WithField("o", "Rename").Error("refuse rename ", args.Path, " to ", args.NewName, ok)
		return &service.RenameReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "Rename").Info("request rename ", args.Path, " to ", args.NewName)
	return &service.RenameReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Delete(ctx context.Context, args *service.DeleteArgs) (*service.DeleteReply, error) {
	ok, blocks := metadatamanager.Delete(args.Path)
	if ok != 0 {
		log.WithField("o", "Delete").Error("refuse delete ", args.Path, ok)
		return &service.DeleteReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "Delete").Info("request delete ", args.Path)
	for _, v := range blocks {
		for _, j := range nn.cache[v] {
			nn.DataNodeList[j].cleanblocks <- v
		}
		delete(nn.cache, v)
	}
	return &service.DeleteReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Stat(ctx context.Context, args *service.StatArgs) (*service.StatReply, error) {
	info, ok := metadatamanager.Stat(args.Path)
	if ok != 0 {
		log.WithField("o", "Stat").Error("refuse stat ", args.Path, ok)
		return &service.StatReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	log.WithField("o", "Stat").Info("request stat ", args.Path)
	return &service.StatReply{Info: info, Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) GetLease(ctx context.Context, args *service.GetLeaseArgs) (*service.GetLeaseReply, error) {
	if nn.Status == 0 {
		return &service.GetLeaseReply{Status: service.StatusCode_NotOK, NameNodeStatus: service.NameNodeStatus_nSafeMode}, nil
	}
	err := nn.lease.Get()
	if err != nil {
		return &service.GetLeaseReply{Status: service.StatusCode_NotOK, NameNodeStatus: service.NameNodeStatus_nWritting}, nil
	}
	return &service.GetLeaseReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) RenewLease(ctx context.Context, args *service.RenewLeaseArgs) (*service.RenewLeaseReply, error) {
	nn.lease.Renew()
	return &service.RenewLeaseReply{}, nil
}

func (nn *NameNode) ReleaseLease(ctx context.Context, args *service.ReleaseLeaseArgs) (*service.ReleaseLeaseReply, error) {
	nn.lease.Release()
	return &service.ReleaseLeaseReply{}, nil
}
