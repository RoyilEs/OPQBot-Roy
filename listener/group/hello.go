package group

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
)

func Hello(ctx context.Context, event events.IEvent) {
	groupMsg := event.ParseGroupMsg()
	if groupMsg.ParseTextMsg().GetTextContent() == "hello" {
		err := apiBuilder.New(global.OBQBotUrl, global.BotQQ).
			SendMsg().GroupMsg().TextMsg("hello  d=====(￣▽￣*)b").ToUin(groupMsg.GetGroupUin()).Do(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}
}
