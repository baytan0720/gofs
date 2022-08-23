package namenode

import (
	"context"
	"gofs/src/namenode/pkg/metamanager"
	"gofs/src/service"
	"time"
)

func (nn *NameNode) GetSystemInfo(ctx context.Context, args *service.GetSystemInfoArgs) (*service.GetSystemInfoReply, error) {
	rep := &service.GetSystemInfoReply{
		NnStatus:       service.NameNodeStatus(nn.status),
		ReplicaNum:     int32(nn.ReplicaNum),
		BlockSize:      nn.BlockSize,
		DataNodes:      make([]*service.DataNodeInfo, 0, nn.DataNodeNum),
		TotalDiskQuota: nn.totaldiskquota,
		UsedDisk:       nn.useddisk,
	}
	for _, v := range nn.DataNodeList {
		if v == nil {
			continue
		}
		rep.DataNodes = append(rep.DataNodes, v.moreinfo)
	}
	return rep, nil
}

func (nn *NameNode) PutFile(ctx context.Context, args *service.PutFileArgs) (*service.PutFileReply, error) {
	rep := &service.PutFileReply{}

	//检查目录树
	check, parentid := metamanager.CheckPath(args.Path, args.FileName)
	if check != service.FileStatus_fPathFound {
		rep.Status = service.StatusCode_NotOK
		rep.FileStatus = check
		return rep, nil
	}

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
	entryid := metamanager.NewEntryId()
	WaitCofirm[entryid] = &tempfile{blockid: blockid, cofirm: make([]int, blockNum), parentid: parentid, size: args.FileSize, filename: args.FileName}

	rep.BlockSize = nn.BlockSize
	return rep, nil
}

// func (nn *NameNode) PutBlock(ctx context.Context, args *service.PutBlockArgs) (*service.PutBlockReply, error) {

// }

// func (nn *NameNode) GetFile(ctx context.Context, args *service.GetFileArgs) (*service.GetFileReply, error) {
// 	return nil, nil
// }

func (nn *NameNode) Mkdir(ctx context.Context, args *service.MkdirArgs) (*service.MkdirReply, error) {
	ok := metamanager.Mkdir(args.Path, args.DirName)
	if ok != 0 {
		return &service.MkdirReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	return &service.MkdirReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) List(ctx context.Context, args *service.ListArgs) (*service.ListReply, error) {
	files, ok := metamanager.List(args.Path)
	if ok != 0 {
		return &service.ListReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	return &service.ListReply{Files: files, Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Rename(ctx context.Context, args *service.RenameArgs) (*service.RenameReply, error) {
	ok := metamanager.Rename(args.Path, args.NewName)
	if ok != 0 {
		return &service.RenameReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	return &service.RenameReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Delete(ctx context.Context, args *service.DeleteArgs) (*service.DeleteReply, error) {
	ok := metamanager.Delete(args.Path)
	if ok != 0 {
		return &service.DeleteReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	return &service.DeleteReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) Stat(ctx context.Context, args *service.StatArgs) (*service.StatReply, error) {
	info, ok := metamanager.Stat(args.Path)
	if ok != 0 {
		return &service.StatReply{Status: service.StatusCode_NotOK, FileStatus: ok}, nil
	}
	return &service.StatReply{Info: info, Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) GetLease(ctx context.Context, args *service.GetLeaseArgs) (*service.GetLeaseReply, error) {
	if nn.status == 0 {
		return &service.GetLeaseReply{Status: service.StatusCode_NotOK, NameNodeStatus: service.NameNodeStatus_nSafeMode}, nil
	}
	if nn.lease != 0 {
		return &service.GetLeaseReply{Status: service.StatusCode_NotOK, NameNodeStatus: service.NameNodeStatus_nWritting}, nil
	}
	nn.lease = -1
	nn.leasetimer.Reset(10 * time.Minute)
	go func() {
		<-nn.leasetimer.C
		nn.lease = 0
	}()
	return &service.GetLeaseReply{Status: service.StatusCode_OK}, nil
}

func (nn *NameNode) RenewLease(ctx context.Context, args *service.RenewLeaseArgs) (*service.RenewLeaseReply, error) {
	nn.leasetimer.Reset(10 * time.Minute)
	return &service.RenewLeaseReply{}, nil
}

func (nn *NameNode) ReleaseLease(ctx context.Context, args *service.ReleaseLeaseArgs) (*service.ReleaseLeaseReply, error) {
	nn.leasetimer.Reset(0)
	tempfile := WaitCofirm[args.EntryId]
	if tempfile.cofirmnum == len(tempfile.blockid) {
		metamanager.Put(tempfile.parentid, tempfile.filename, args.EntryId, tempfile.size, time.Now().Format(time.Now().Format("2006-01-02 15:04:05")), tempfile.blockid)
	}
	WaitCofirm[args.EntryId] = nil
	return &service.ReleaseLeaseReply{}, nil
}
