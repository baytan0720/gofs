package model

import (
	"fmt"
	"gofs/DataNode/internal/service"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"
)

//ListenSocket 用于建立socket连接，等待文件传输
func (dn *DataNode) ListenSocket() {
	server, err := net.Listen("tcp", dn.Addr)
	log.Println(dn.Addr + " is Listening")
	if err != nil {
		log.Println(err)
	}
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
		}
		fmt.Println("client connected")
		//bufio.NewReader(connection)
		go processClient(connection, dn)
	}
}

func processClient(conn net.Conn, dn *DataNode) {
	buffer := make([]byte, 2048)
	err := conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		log.Println(err)
	}
	mLen, err := conn.Read(buffer)
	if err != nil || mLen == 0 {
		fmt.Println("Error reading:", err.Error())
	}
	req := &service.ReplicationRequest{}
	err = proto.Unmarshal(buffer[:mLen], req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Receive code ", req.OperationType)
	if req.OperationType == 2 {
		// 表示客户端发送Get请求
		file, _ := ReadDisk(req.BlockName, "7515")
		res := &service.ReplicationResponse{
			ResponseType: 2,
			File:         file,
		}
		d, _ := proto.Marshal(res)
		conn.Write(d)
		fmt.Println("Write Done ", res.ResponseType)
		// 等待Write完毕且客户端Read完毕
		time.Sleep(time.Second)

	} else {
		// 表示客户端发送Put请求
		// 循环 直到OT==0表示block上传完毕
		for req.OperationType == 1 {
			// block落盘
			block, err := WriteDisk(req.File, strings.Split(dn.Addr, ":")[1])
			dn.Blocklist = append(dn.Blocklist, block)

			// 写入反馈
			res := &service.ReplicationResponse{
				ResponseType: 1,
				Error:        nil,
				BlockName:    block.Name,
			}
			if err != nil {
				res.ResponseType = 0
				res.Error = append(res.Error, &service.ReplicationError{ErrorContent: err.Error()})
			}
			d, _ := proto.Marshal(res)
			conn.Write(d)

			// 读取Client
			buffer = make([]byte, 2048)
			err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			if err != nil {
				log.Println(err)
			}
			mLen, err = conn.Read(buffer)
			if err != nil || mLen == 0 {
				fmt.Println("Error reading:", err.Error())
			}
			req = &service.ReplicationRequest{}
			err = proto.Unmarshal(buffer[:mLen], req)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("Receive code ", req.OperationType)
		}
	}
	log.Println(conn.RemoteAddr().String() + " closed")
	conn.Close()
}
