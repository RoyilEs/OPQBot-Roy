package main

import (
	"github.com/charmbracelet/log"
	"obqbot/global"
	"obqbot/models/pixiv"
)

func main() {
	test, _ := pixiv.NewPixivTest()
	log.Info(test)
	response, err := pixiv.NewPixiv().Set().DoPixiv()
	if err != nil {
		panic(err)
	}
	global.Log.Info(response)

}
