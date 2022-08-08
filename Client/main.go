package main

import (
	"fmt"
	"gofs/Client/service"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	_ "net/http/pprof"
	"net/rpc"
	"os"
)

func main() {
	path := "C:\\Users\\HP\\Desktop\\a.txt"

	//put(path)
	get(path, "")
}

type Chunk struct {
	Sequence  int
	BlockName string
	Addr      []string
}

func get(path, remotePath string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	// TODO 访问NN 获取目标文件所有block信息和block位置
	addr := []string{":10745", ":."}
	fileMap := []Chunk{
		{
			Sequence:  1,
			BlockName: "6e70f19c1d7f7b23",
			Addr:      addr,
		},
		{
			Sequence:  2,
			BlockName: "955daacb0bb2509d",
			Addr:      addr,
		},
		{
			Sequence:  3,
			BlockName: "412197a8926479ee",
			Addr:      addr,
		},
		{
			Sequence:  4,
			BlockName: "7645771a4f357794",
			Addr:      addr,
		},
		{
			Sequence:  5,
			BlockName: "f383534614a54039",
			Addr:      addr,
		},
	}
	for i := 0; i < len(fileMap); i++ {
		con, err := net.Dial("tcp", fileMap[i].Addr[0])
		if err != nil {
			log.Println(err)
		}
		req := &service.ReplicationRequest{
			OperationType: 2,
			BlockName:     fileMap[i].BlockName,
		}
		d, _ := proto.Marshal(req)
		con.Write(d)

		// 读取回复的文件数据
		b := make([]byte, 2048)
		ii, _ := con.Read(b)
		res := &service.ReplicationResponse{}
		err = proto.Unmarshal(b[:ii], res)
		if err != nil {
			log.Println(err)
		}
		if res.ResponseType == 2 {
			ii, _ = file.Write(res.File.Payload)
			log.Printf("Sequence %d has done \n", i+1)
			// TODO 需要判断。如果 i!=res.File.Length 重写
		} else {
			log.Println("Error ResponseType should be 2, actually is ", res.String())
		}
		con.Close()
	}
}

func put(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	const ChunkSize = 1 * (1 << 10) // 1kb
	info, _ := file.Stat()
	// 计算需要分割的份数
	blockNum := int(math.Ceil(float64(info.Size()) / float64(ChunkSize)))
	fmt.Printf("Split file into %d pieces\n", blockNum)
	// 以下需要保证客户端把所有block传输完毕且传输成功才能退出
	for i := 0; i < blockNum; i++ {
		blockSize := int(math.Min(ChunkSize, float64(info.Size()-int64(ChunkSize*i))))
		bytes := make([]byte, blockSize)
		file.Read(bytes)
		log.Printf("Read File Block [%d] %d bit \n", i, blockSize)

		// TODO 访问NN，获取目标端口
		ports := []string{":7515"}

		// 建立TCP连接
		con, err := net.Dial("tcp", ports[0])
		if err != nil {
			log.Println(err)
		}
		sendLoop(con, bytes, i, false)
		// receiveLoop接收服务都发送的反馈信息，应该是阻塞的
		res := receiveLoop(con)

		// ResponseType == 0 表示写入错误，自旋重试
		for res.ResponseType == 0 {
			errors := res.Error
			for e := range errors {
				log.Println(e)
			}
			log.Println("Rewrite block ", res.BlockName)
			sendLoop(con, bytes, i, false)
		}
		// 表示此次文件上传完毕
		sendLoop(con, nil, 0, true)
		con.Close()
	}

}

func sendLoop(w io.Writer, content []byte, index int, lastOne bool) {
	req := &service.ReplicationRequest{
		OperationType: 0,
		File:          nil,
	}
	if !lastOne {
		replicationFile := &service.ReplicationFile{
			Payload:  content,
			Length:   uint64(len(content)),
			Sequence: uint32(index + 1),
		}
		req.File = replicationFile
		req.OperationType = 1
	}
	d, _ := proto.Marshal(req)
	w.Write(d)
	if !lastOne {
		fmt.Println("Write Done ", req.File.Sequence)
	}

}

func receiveLoop(r io.Reader) *service.ReplicationResponse {
	readByte := make([]byte, 1024)
	i, _ := r.Read(readByte)
	res := &service.ReplicationResponse{}
	err := proto.Unmarshal(readByte[:i], res)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Decode ", res.String())
	return res
}

func Hello(nn *rpc.Client) {
	Args := &HelloArgs{}
	Reply := &HelloReply{}
	err := nn.Call("NameNode.Hello", &Args, &Reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(Reply.S)
}

func register() *rpc.Client {
	nn, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Connect fail: ", err)
		os.Exit(0)
	}
	return nn
}

func logo() {
	fmt.Println("  _____  ____  ______ _____ ")
	fmt.Println(" / ____|/ __ \\|  ____/ ____|")
	fmt.Println("| |  __| |  | | |__ | (___  ")
	fmt.Println("| | |_ | |  | |  __| \\___ \\ ")
	fmt.Println("| |__| | |__| | |    ____) |")
	fmt.Println(" \\_____|\\____/|_|   |_____/ ")
	fmt.Println("")
}
