package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"obqbot/config"
	"obqbot/global"
)

const ConfigFile = "settings.yaml"

// InitConf 初始化读取配置文件
func InitConf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("yamlFile.Get err   #%v ", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("yaml.Unmarshal: %v", err)
	}
	log.Println("config yamlFile InitConf success")
	global.Config = c
}
