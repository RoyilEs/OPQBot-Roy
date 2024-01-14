package friend

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
)

func Hello(ctx context.Context, event events.IEvent) {
	groupMsg := event.ParseFriendMsg()
	if groupMsg.ParseTextMsg().GetTextContent() == "hello" {
		err := apiBuilder.New(global.OBQBotUrl, global.BotQQ).
			SendMsg().FriendMsg().TextMsg("你好你好").ToUin(groupMsg.GetFriendUin()).Do(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}
}
