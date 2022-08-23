package main

import "gofs/src/namenode/pkg/namenode"

func main() {
	nn := namenode.MakeNameNode()
	nn.Server()
}
