package main

import (
	"gofs/NameNode/internal/model"
)

func main() {
	nn := model.MakeNameNode()
	nn.Server()
}

/*//rpc调用示例
func (nn *model.NameNode) Hello(Args *HelloArgs, Reply *HelloReply) error {
	Reply.S = "Hello"
	return nil
}*/

// func (nn *NameNode) close() {

// }
