package main

import (
	"fmt"
	"gofs/Client/config"
	"gofs/Client/internal/model"
	"os"
)

func main() {
	config.Opencfg()
	if len(os.Args) < 2 {
		fmt.Println("input 'gofs help' to get usage")
		return
	}
	switch os.Args[1] {
	case "dninfo":
		model.DNInfo()
	case "help":
		model.Help()
	case "put":
		if len(os.Args) != 4 {
			fmt.Println("Invalid Argument, Example: gofs put path1 path2")
			return
		}
		model.Put(os.Args[2], os.Args[3])
	default:
		fmt.Println("input 'gofs help' to get usage")
	}
}
