package main

import (
	"gofs/NameNode/internal/model"
)

func main() {
	nn := model.MakeNameNode()
	nn.Server()
}
