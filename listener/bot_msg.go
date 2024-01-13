package listener

import (
	"context"
	"encoding/json"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	Ok3Http "obqbot/utils"
	"strconv"
)

type YiResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    Data   `json:"data"`
}

type Data struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
}

func Yian(ctx context.Context, event events.IEvent) {
	groupMsg := event.ParseGroupMsg()
	if groupMsg.ParseTextMsg().GetTextContent() == "一言" {
		s := Ok3Http.NewHTTPClient("http://www.wudada.online/Api/MrYy")
		get, err := s.DoGet("", nil)
		var data YiResponse
		if err != nil {
			global.Log.Error(err)
			return
		}
		if err != nil {
			err := apiBuilder.New(global.OBQBotUrl, global.BotQQ).SendMsg().
				GroupMsg().TextMsg("获取一言失败").ToUin(groupMsg.GetGroupUin()).Do(ctx)
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
		global.Log.Info(data.Data.Hitokoto + "\t——" + data.Data.From + ":" + strconv.FormatInt(groupMsg.GetGroupUin(), 10))
		err = apiBuilder.New(global.OBQBotUrl, global.BotQQ).SendMsg().
			GroupMsg().TextMsg(data.Data.Hitokoto + "\t——" + data.Data.From).ToUin(groupMsg.GetGroupUin()).Do(ctx)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}
}
