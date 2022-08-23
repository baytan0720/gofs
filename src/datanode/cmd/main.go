package main

import "gofs/src/datanode/pkg/datanode"

func main() {
	dn := datanode.MakeDataNode()
	dn.Server()
}
