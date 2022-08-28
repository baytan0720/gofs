package main

import (
	"fmt"
	"gofs/src/client/pkg/api"
	"log"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

type config struct {
	NameNodeAddr string
	NameNodePort string
}

var Config *config

func main() {
	if len(os.Args) < 2 {
		fmt.Println("input 'gofs help' to get usage")
		return
	}
	if os.Args[1] == "help" {
		api.Help()
		return
	}
	Config = &config{}
	opencfg()
	api.Addr = Config.NameNodeAddr + Config.NameNodePort
	switch os.Args[1] {
	case "sysinfo":
		api.SysInfo()
	case "put":
		if len(os.Args) != 4 {
			fmt.Println("Invalid Argument, usage: 'gofs put <gofspath> <localpath>'")
			return
		}
		api.Put(os.Args[2], os.Args[3])
	case "get":
		if len(os.Args) != 4 {
			fmt.Println("Invalid Argument, usage: 'gofs get <gofspath> <localpath>'")
			return
		}
		api.Get(os.Args[2], os.Args[3])
	case "ls":
		if len(os.Args) != 3 {
			fmt.Println("Invalid Argument, usage: 'gofs ls <gofspath>'")
			return
		}
		api.List(os.Args[2])
	case "mkdir":
		if len(os.Args) != 4 {
			fmt.Println("Invalid Argument, usage: 'gofs mkdir <gofspath> <dirname>'")
			return
		}
		api.Mkdir(os.Args[2], os.Args[3])
	case "rename":
		if len(os.Args) != 4 {
			fmt.Println("Invalid Argument, usage: 'gofs rename <gofspath>'")
			return
		}
		api.Rename(os.Args[2], os.Args[3])
	case "stat":
		if len(os.Args) != 3 {
			fmt.Println("Invalid Argument, usage: 'gofs stat <gofspath>'")
			return
		}
		api.Stat(os.Args[2])
	case "del":
		if len(os.Args) != 3 {
			fmt.Println("Invalid Argument, usage: 'gofs del <gofspath>'")
			return
		}
		api.Delete(os.Args[2])
	default:
		fmt.Println("usage: 'gofs help' to get usage of gofs")
	}
}

func opencfg() {
	var path string
	if runtime.GOOS == "windows" {
		pwd, _ := os.Getwd()
		path = pwd + ""
	} else {
		path = "../../../config/config.toml"
	}
	_, err := toml.DecodeFile(path, Config)
	if err != nil {
		log.Fatal("Config Read Fail: ", err)
	}
}
