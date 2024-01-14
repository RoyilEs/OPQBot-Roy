package group

import (
	"context"
	"encoding/json"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	Ok3Http "obqbot/utils"
)

type data struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Res  result `json:"result"`
}

type result struct {
	Content string `json:"content"`
}

func GoodNight(ctx context.Context, event events.IEvent) {
	groupMsg := event.ParseGroupMsg()
	apiBuilderNew := global.ApiBuilderNew
	if groupMsg.ParseTextMsg().GetTextContent() == "晚安" || groupMsg.ParseTextMsg().GetTextContent() == "gn" {
		s := Ok3Http.NewHTTPClient("https://apis.tianapi.com/wanan/index?key=35cece84b9177c32006439a30b95c57e")
		get, err := s.DoGet("", nil)
		var data data
		if err != nil {
			global.Log.Error(err)
			return
		}
		if err != nil {
			err := apiBuilderNew.SendMsg().
				GroupMsg().TextMsg("获取晚安语句失败").ToUin(groupMsg.GetGroupUin()).Do(ctx)
			if err != nil {
				return
			}
			return
		}
		err = json.Unmarshal(get, &data)
		if err != nil {
			global.Log.Error(err)
			return
		}
		err = apiBuilderNew.SendMsg().
			GroupMsg().TextMsg(data.Res.Content).ToUin(groupMsg.GetGroupUin()).
			At(groupMsg.GetSenderUin()).
			Do(ctx)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}
}
