package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//结构体字段首字母必须大写
type config struct {
	NumDataNodeLimit int
	Port             string
}

var Config *config

func Opencfg() {
	pwd, _ := os.Getwd()
	path := pwd + "\\NameNode\\config\\config.toml"
	Config = &config{}
	_, err := toml.DecodeFile(path, Config)
	if err != nil {
		log.Fatal("Config read fail ", err.Error())
	}
}
