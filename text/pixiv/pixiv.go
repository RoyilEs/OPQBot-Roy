package main

import (
	"github.com/charmbracelet/log"
	"obqbot/models/pixiv"
)

func main() {
	size, _ := pixiv.NewPixiv().Set().DoQuery()
	pixivResponse, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, size)
	log.Info(pixivResponse)

}
