package model

import (
	"context"
	"gofs/NameNode/config"
	"gofs/NameNode/internal/service"
	"strconv"

	"github.com/yitter/idgenerator-go/idgen"
)

//请求DataNode状态信息
func (nn *NameNode) DNInfo(ctx context.Context, args *service.DNInfoArgs) (*service.DNInfoReply, error) {
	dnlist := make([]*service.DNInfo, 0, len(nn.DataNodeList))
	for _, v := range nn.DataNodeList {
		if v != nil && v.alive == 1 {
			blocklist := make([]string, 0, len(v.Blocklist))
			for _, v := range v.Blocklist {
				blocklist = append(blocklist, v.Id)
			}
			dnlist = append(dnlist, &service.DNInfo{Id: uint32(v.Id), BlockList: blocklist})
		}
	}
	return &service.DNInfoReply{DNList: dnlist}, nil
}

//请求上传文件
func (nn *NameNode) PutFile(ctx context.Context, args *service.PutFileArgs) (*service.PutFileReply, error) {
	//检查目录树

	//计算block数目
	blocksize := config.Config.BlockSize << 20
	splitsum := args.Size / (blocksize)
	if args.Size%(blocksize) != 0 {
		splitsum += 1
	}

	return &service.PutFileReply{ACK: 1, BlockSize: config.Config.BlockSize, BlockId: getBlockId(int(splitsum))}, nil
}

func (nn *NameNode) PutBlock(ctx context.Context, args *service.PutBlockArgs) (*service.PutBlockReply, error) {
	//选择接受上传的DataNode
	dnlist := make([]*service.DNAddr, 3)
	count := 0
	for i := 0; i < len(nn.DataNodeList); i++ {
		if count == 3 {
			break
		}
		if nn.DataNodeList[i] != nil && nn.DataNodeList[i].alive == 1 {
			dnlist[count] = &service.DNAddr{Id: uint32(nn.DataNodeList[i].Id)}
			count++
		}
	}
	if count < 3 {
		return &service.PutBlockReply{ACK: 0, Err: "Not enough DataNode"}, nil
	}
	return &service.PutBlockReply{ACK: 1, DNList: dnlist}, nil
}

func getBlockId(sum int) []string {
	blockid := make([]string, sum)
	for i := 0; i < sum; i++ {
		blockid[i] = strconv.FormatInt(idgen.NextId(), 10) + "_" + strconv.Itoa(i)
	}
	return blockid
}
