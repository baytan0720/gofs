package namenode

import (
	"context"
	"gofs/src/service"
	"log"
	"strconv"
	"time"

	"github.com/yitter/idgenerator-go/idgen"
)

type tempfile struct {
	filename  string
	parentid  int64
	size      int64
	blockid   []string
	cofirm    []int
	cofirmnum int
}

type DataNode struct {
	dietimer *time.Timer
	moreinfo *service.DataNodeInfo
}

var WaitCofirm map[int64]*tempfile = make(map[int64]*tempfile)

func (nn *NameNode) Register(ctx context.Context, args *service.RegisterArgs) (*service.RegisterReply, error) {
	var id int32
	select {
	case id = <-nn.idChan:
	default:
		nn.putId()
		id = <-nn.idChan
	}

	timer := time.NewTimer(time.Duration(nn.HeartBeatTimeout) * time.Second)
	newDn := &DataNode{
		dietimer: timer,
		moreinfo: args.Info,
	}
	newDn.moreinfo.Id = id
	newDn.moreinfo.Port = ":" + strconv.Itoa(int(id)+1024)
	nn.DataNodeList[id] = newDn

	go func() {
		<-timer.C
		nn.totaldiskquota -= nn.DataNodeList[id].moreinfo.DiskQuota
		nn.useddisk -= nn.DataNodeList[id].moreinfo.UsedDisk
		nn.DataNodeList[id] = nil
		log.Println("ID: ", id, " is died")
		nn.DataNodeNum--
		if nn.DataNodeNum < 3 && nn.status == 1 {
			nn.status = 0
		}
		nn.idChan <- id
	}()

	rep := &service.RegisterReply{Status: service.StatusCode_OK, Id: id}
	log.Println("ID: ", rep.Id, " is connected")
	nn.DataNodeNum++
	nn.totaldiskquota += newDn.moreinfo.DiskQuota
	nn.useddisk += newDn.moreinfo.UsedDisk
	if nn.DataNodeNum >= 3 && nn.status == 0 {
		nn.status = 1
	}
	return rep, nil
}

func (nn *NameNode) HeartBeat(ctx context.Context, args *service.HeartBeatArgs) (*service.HeartBeatReply, error) {
	if nn.DataNodeList[args.Id] == nil {
		rep := &service.HeartBeatReply{Status: service.StatusCode_NotRegister}
		return rep, nil
	}
	nn.DataNodeList[args.Id].dietimer.Stop()
	nn.DataNodeList[args.Id].dietimer.Reset(time.Duration(nn.HeartBeatTimeout) * time.Second)
	rep := &service.HeartBeatReply{Status: service.StatusCode_OK}
	return rep, nil
}

func (nn *NameNode) BlockReport(ctx context.Context, args *service.BlockReportArgs) (*service.BlockReportReply, error) {
	if nn.DataNodeList[args.Id] == nil {
		rep := &service.BlockReportReply{Status: service.StatusCode_NotRegister}
		return rep, nil
	}
	nn.DataNodeList[args.Id].moreinfo.Blocks = args.Blocks
	nn.DataNodeList[args.Id].moreinfo.BlockNum = int32(len(args.Blocks))
	rep := &service.BlockReportReply{Status: service.StatusCode_OK}
	return rep, nil
}

func (nn *NameNode) NewBlockReport(ctx context.Context, args *service.NewBlockReportArgs) (*service.NewBlockReportReply, error) {
	nn.DataNodeList[args.Id].moreinfo.Blocks = append(nn.DataNodeList[args.Id].moreinfo.Blocks, args.Block)
	nn.DataNodeList[args.Id].moreinfo.BlockNum++
	nn.DataNodeList[args.Id].moreinfo.UsedDisk += args.Block.Size
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
	rep := &service.NewBlockReportReply{Status: service.StatusCode_OK}
	return rep, nil
}

func (nn *NameNode) putId() {
	for i := 0; i < 3; i++ {
		nn.idChan <- nn.idIncrease
		nn.idIncrease++
	}
	if int(nn.idIncrease) > len(nn.DataNodeList)-1 {
		nn.DataNodeList = append(nn.DataNodeList, make([]*DataNode, len(nn.DataNodeList))...)
	}
}

func getBlockId(sum int) []string {
	blockid := make([]string, sum)
	for i := 0; i < sum; i++ {
		blockid[i] = strconv.FormatInt(idgen.NextId(), 10) + "^" + strconv.Itoa(i)
	}
	return blockid
}
