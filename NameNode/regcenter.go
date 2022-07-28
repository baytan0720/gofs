package main

import "log"

func (nn *NameNode) DNRegister(Arg *Args, Reply *Reply) error {
	nn.NumDataNode++
	log.Println("A DataNode register sucess, NumDataNode: ", nn.NumDataNode)
	return nil
}
