package config

import (
	"log"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
)

//结构体字段首字母必须大写
type config struct {
	Addr string
	Port string
	GOOS string
}

var Config *config

func Opencfg() {
	var path string
	sysType := runtime.GOOS
	if sysType == "windows" {
		pwd, _ := os.Getwd()
		path = pwd + "\\Client\\config\\config.toml"
	} else {
		path = "../config/config.toml"
	}
	Config = &config{GOOS: sysType}
	_, err := toml.DecodeFile(path, Config)
	if err != nil {
		log.Fatal("Config read fail ", err.Error())
	}
}