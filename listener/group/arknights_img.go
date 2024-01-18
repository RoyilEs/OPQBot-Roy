package group

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"golang.org/x/time/rate"
	"obqbot/common/limiter"
	"obqbot/global"
	"obqbot/models/pixiv"
	"obqbot/utils"
	"strings"
	"time"
)

func ArknightsImg(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		var ZhouTest = []string{"来点粥图", "来张粥图", "粥图一张"}

		query, err := pixiv.NewPixiv().Set().DoQuery()
		if err != nil {
			global.Log.Error(err)
			return
		}
		iPixiv, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)

		if utils.IsInListToS(text, ZhouTest) {
			query, err := pixiv.NewPixiv().Set().DoQuery()
			if err != nil {
				global.Log.Error(err)
				return
			}
			iPixiv, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				buf := i.UrlToBase64(pixiv.ModifyPixivImageUrl(url))
				//TODO 已修改图片质量需要所有修改的钩子函数并作接口化处理
				pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
					GroupPic().SetBase64Buf(buf).DoUpload(ctx)
				if err != nil {
					global.Log.Error(err)
					return
				}
				log.Debug(pic)
				log.Info(url)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).TextMsg("原图链接:" + url).Do(ctx)
			}
		}

		var ZhouTest2 = []string{"明日方舟", "来点舟图", "舟图一张"}
		if utils.IsInListToS(text, ZhouTest2) {
			for _, i := range iPixiv.GetData() {
				url := i.GetDataUrls().GetSize()
				pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
					GroupPic().SetFileUrlPath(url).DoUpload(ctx)
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

func PixivImg(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		limit := limiter.NewLimiter(rate.Every(1*time.Second), 2, "")
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		split := strings.Split(text, " ")
		// 令牌桶限流器--防止大量请求
		if !limit.Allow() {
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("请求过于频繁").Do(ctx)
			return
		}
		if split[0] == "cnm" {
			query, err := pixiv.NewPixiv().Set().SetTag(split[1]).DoQuery()
			if err != nil {
				global.Log.Error(err)
				return
			}
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("下载速度比较慢 请耐心等待").Do(ctx)
			iPixiv, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)
			for _, data := range iPixiv.GetData() {
				url := data.GetDataUrls().GetSize()
				pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
					GroupPic().SetFileUrlPath(url).DoUpload(ctx)
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
