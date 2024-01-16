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
			iPixiv := pixiv.NewPixiv().GetData()

			for _, i := range iPixiv {
				url := i.GetDataUrls().GetThumb()
				buf := i.DoThumbToBase64(url)
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

				//apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				//	GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("原图链接:" + pixiv.ModifyPixivImageUrl(url)).Do(ctx)
			}
		}
	}
}
