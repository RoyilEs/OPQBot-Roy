package group

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	"obqbot/models/pixiv"
	"obqbot/utils"
)

func PixivMsg(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		var ZhouTest = []string{"来点粥图", "来张粥图", "粥图一张"}
		if utils.IsInListToS(text, ZhouTest) {
			iPixiv, _ := pixiv.NewPixivTest()
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				buf := i.UrlToBase64(url)
				pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
					GroupPic().SetBase64Buf(buf).DoUpload(ctx)
				if err != nil {
					global.Log.Error(err)
					return
				}
				log.Debug(pic)
				log.Info(url)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).TextMsg("原图链接:" + pixiv.ModifyPixivImageUrl(url)).Do(ctx)
			}
		}

		var ZhouTest2 = []string{"明日方舟", "来点舟图", "舟图一张"}
		if utils.IsInListToS(text, ZhouTest2) {
			iPixiv, _ := pixiv.NewPixivTest()
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				modifiedUrl := pixiv.ModifyPixivImageUrl(url)
				pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
					GroupPic().SetFileUrlPath(modifiedUrl).DoUpload(ctx)
				if err != nil {
					global.Log.Error(err)
					return
				}
				log.Debug(pic)
				log.Info(url)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).Do(ctx)
			}
		}
	}
}
