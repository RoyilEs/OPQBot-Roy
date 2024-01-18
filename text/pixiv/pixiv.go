package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"obqbot/models/pixiv"
)

func main() {
	size, _ := pixiv.NewPixiv().Set().DoQuery()
	pixivResponse, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, size)
	log.Info(pixivResponse)
	getSize := pixivResponse.GetData()[0].GetDataUrls().GetSize()

	fmt.Println(pixiv.ModifyPixivImageUrl(getSize))
}
