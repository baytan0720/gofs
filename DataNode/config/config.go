package config

import (
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

//结构体字段首字母必须大写
type config struct {
	Addr    string
	Port    string
	GOOS    string
	LocalIP string
}

var Config *config

func Opencfg() {
	var path string
	sysType := runtime.GOOS
	if sysType == "windows" {
		pwd, _ := os.Getwd()
		path = pwd + "\\DataNode\\config\\config.toml"
	} else {
		path = "../config/config.toml"
	}
	Config = &config{GOOS: sysType}
	_, err := toml.DecodeFile(path, Config)
	if err != nil {
		log.Fatal("Config read fail ", err)
	}
	Config.LocalIP = GetLocalIP()
}

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println("check port :53")
	}
	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
	conn.Close()
	return ip
}
