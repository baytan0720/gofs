package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	nn := register()

	for {
		var cmd string
		fmt.Print("gofs > ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "exit":
			return
		case "hello":
			Hello(nn)
		case "help":
		default:
			fmt.Println("Unknow command: You could press \"help\" to get all commands.")
		}
	}
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
