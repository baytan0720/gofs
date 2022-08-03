package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

//结构体字段首字母必须大写
type config struct {
	Addr string
	Port string
}

var Config *config

func opencfg() {
	path := "./config/config.toml"
	Config = &config{}
	_, err := toml.DecodeFile(path, Config)
	if err != nil {
		log.Fatal("Config read fail")
	}
}
