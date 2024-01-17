package group

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	"time"
)

func Img(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ExcludeAtInfo().ParseTextMsg().GetTextContent()
		if groupMsg.AtBot() {
			if text == "test" {
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("test").DoWithCallBack(ctx, func(iApiBuilder *apiBuilder.Response, err error) {
					response, err := iApiBuilder.GetGroupMessageResponse()
					if err != nil {
						return
					}
					time.Sleep(time.Second * 1)
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).GroupManager().RevokeMsg().ToGUin(groupMsg.GetGroupUin()).MsgSeq(response.MsgSeq).MsgRandom(response.MsgTime).Do(ctx)
				})
			}
		}
		if text == "img" {
			//client := Ok3Http.NewHTTPClient("https://i.pixiv.re/img-original/img/2020/08/27/03/29/53/83959163_p1.jpg")
			////body, err := client.DoGet("", nil)
			//if err != nil {
			//	panic(err)
			//}
			pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().GroupPic().SetFileUrlPath("https://i.pixiv.re/img-original/img/2024/01/12/03/01/42/115092758_p0.jpg").DoUpload(ctx)
			if err != nil {
				panic(err)
			}
			log.Debug(pic)
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).Do(ctx)
		}
	}

}
