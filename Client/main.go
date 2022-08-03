package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func main() {
	nn := register()

	reader := bufio.NewReader(os.Stdin)
	var cmd []byte
	fmt.Println("  _____  ____  ______ _____ ")
	fmt.Println(" / ____|/ __ \\|  ____/ ____|")
	fmt.Println("| |  __| |  | | |__ | (___  ")
	fmt.Println("| | |_ | |  | |  __| \\___ \\ ")
	fmt.Println("| |__| | |__| | |    ____) |")
	fmt.Println(" \\_____|\\____/|_|   |_____/ ")
	fmt.Println("")
	for {
		fmt.Print("gofs > ")
		cmd, _, _ = reader.ReadLine()
		switch string(cmd) {
		case "exit":
			return
		case "hello":
			Hello(nn)
		case "help":
			fmt.Println("Usage:")
			fmt.Println("  hello  (Verify the feasibility of the RPC)")
			fmt.Println("  exit")
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
