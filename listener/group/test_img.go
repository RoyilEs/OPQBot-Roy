package group

import (
	"context"
	"encoding/base64"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	Ok3Http "obqbot/utils"
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
			client := Ok3Http.NewHTTPClient("http://8.141.1.249/uploads/file/(5)[IP[MA20W6]S[A]~DG7G.PNG")
			body, err := client.DoGet("", nil)
			if err != nil {
				panic(err)
			}
			pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().GroupPic().SetBase64Buf(base64.StdEncoding.EncodeToString(body)).DoUpload(ctx)
			if err != nil {
				panic(err)
			}
			log.Debug(pic)
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).Do(ctx)
		}
	}

}
