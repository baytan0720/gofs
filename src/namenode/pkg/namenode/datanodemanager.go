package namenode

import (
	"context"
	"gofs/src/namenode/pkg/metadatamanager"
	"gofs/src/service"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/yitter/idgenerator-go/idgen"
)

type datanode struct {
	dietimer    *time.Timer
	cleanblocks chan string
	info        *service.DataNodeInfo
	load        *service.DataNodeLoad
}

func (nn *NameNode) Register(ctx context.Context, args *service.RegisterArgs) (*service.RegisterReply, error) {
	//分配id
	var id int32
	select {
	case id = <-nn.idChan:
	default:
		nn.putId()
		id = <-nn.idChan
	}

	//心跳计时器
	timer := time.NewTimer(time.Duration(nn.HeartBeatTimeout) * time.Second)
	newDn := &datanode{
		dietimer:    timer,
		info:        args.Info,
		cleanblocks: make(chan string, 10),
	}
	newDn.info.Id = id
	newDn.info.Port = ":" + strconv.Itoa(int(id)+1024)
	nn.DataNodeList[id] = newDn

	//心跳超时
	go func() {
		<-timer.C
		for _, v := range nn.DataNodeList[id].info.Blocks {
			delete(nn.cache, v.Id)
		}
		close(nn.DataNodeList[id].cleanblocks)
		nn.DataNodeList[id] = nil
		log.Info("ID: ", id, " is offline")
		nn.DataNodeNum--
		if nn.DataNodeNum < 3 && nn.Status == 1 {
			nn.Status = 0
		}
		nn.idChan <- id
	}()

	rep := &service.RegisterReply{Status: service.StatusCode_OK, Id: id}
	log.Info("ID: ", id, " is online")

	nn.DataNodeNum++
	if nn.DataNodeNum >= 3 && nn.Status == 0 {
		nn.Status = 1
	}
	return rep, nil
}

func (nn *NameNode) HeartBeat(ctx context.Context, args *service.HeartBeatArgs) (*service.HeartBeatReply, error) {
	dn := nn.DataNodeList[args.Id]
	if dn == nil {
		rep := &service.HeartBeatReply{Status: service.StatusCode_NotRegister}
		return rep, nil
	}
	dn.load = args.Load
	dn.dietimer.Stop()
	dn.dietimer.Reset(time.Duration(nn.HeartBeatTimeout) * time.Second)
	rep := &service.HeartBeatReply{Status: service.StatusCode_OK}
	for {
		select {
		case blockid := <-dn.cleanblocks:
			rep.CleanBlockId = append(rep.CleanBlockId, blockid)
		default:
			goto Loop
		}
	}
Loop:
	return rep, nil
}

func (nn *NameNode) BlockReport(ctx context.Context, args *service.BlockReportArgs) (*service.BlockReportReply, error) {
	if nn.DataNodeList[args.Id] == nil {
		rep := &service.BlockReportReply{Status: service.StatusCode_NotRegister}
		return rep, nil
	}
	nn.DataNodeList[args.Id].info.Blocks = args.Blocks
	for _, v := range args.Blocks {
		nn.cache[v.Id] = append(nn.cache[v.Id], int(args.Id))
	}
	rep := &service.BlockReportReply{Status: service.StatusCode_OK}
	return rep, nil
}

func (nn *NameNode) NewBlockReport(ctx context.Context, args *service.NewBlockReportArgs) (*service.NewBlockReportReply, error) {
	nn.DataNodeList[args.Id].info.Blocks = append(nn.DataNodeList[args.Id].info.Blocks, args.Block)
	nn.cache[args.Block.Id] = append(nn.cache[args.Block.Id], int(args.Id))
	tempfile := WaitCofirm[args.EntryId]
	if tempfile == nil {
		return &service.NewBlockReportReply{Status: service.StatusCode_OK}, nil
	}
	for i, v := range tempfile.blockid {
		if v == args.Block.Id && tempfile.cofirm[i] == 0 {
			tempfile.cofirm[i] = 1
			tempfile.cofirmnum++
			break
		}
	}
	log.Info("ID: ", args.Id, " Report a new block: ", args.Block.Id)
	if tempfile.cofirmnum == len(tempfile.blockid) {
		metadatamanager.Put(tempfile.parentid, tempfile.filename, args.EntryId, tempfile.size, time.Now().Format("2006-01-02 15:04:05"), tempfile.blockid)
		log.WithField("o", "PUT").Info(tempfile.filename, " put sucess")
	} else {
		log.WithField("o", "PUT").Error(tempfile.filename, " put fail")
	}
	delete(WaitCofirm, args.EntryId)
	rep := &service.NewBlockReportReply{Status: service.StatusCode_OK}
	return rep, nil
}

func (nn *NameNode) putId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
	if int(nn.idIncrease) > len(nn.DataNodeList)-1 {
		nn.DataNodeList = append(nn.DataNodeList, make([]*datanode, len(nn.DataNodeList))...)
	}
}

func getBlockId(sum int) []string {
	blockid := make([]string, sum)
	for i := 0; i < sum; i++ {
		blockid[i] = strconv.FormatInt(idgen.NextId(), 10)
	}
	return blockid
}
