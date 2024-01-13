package main

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	core2 "obqbot/core"
	"obqbot/global"
	"obqbot/listener"
)

func main() {
	// 读取配置文件
	core2.InitConf()
	// 初始化日志
	global.Log = core2.InitLogger()

	core, err := OPQBot.NewCore(global.OBQBotUrl)
	if err != nil {
		panic(err)
	}

	core.On(events.EventNameGroupMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.ParseGroupMsg()
		if groupMsg.ParseTextMsg().GetTextContent() == "hello" {
			err := apiBuilder.New(global.OBQBotUrl, global.BotQQ).
				SendMsg().GroupMsg().TextMsg("hellod=====(￣▽￣*)b").ToUin(groupMsg.GetGroupUin()).Do(ctx)
			if err != nil {
				log.Error(err)
				return
			}
		}
	})

	core.On(events.EventNameGroupMsg, listener.Yian)

	core.On(events.EventNameFriendMsg, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.ParseFriendMsg()
		if groupMsg.ParseTextMsg().GetTextContent() == "hello" {
			err := apiBuilder.New(global.OBQBotUrl, global.BotQQ).
				SendMsg().FriendMsg().TextMsg("hellod=====(￣▽￣*)b").ToUin(groupMsg.GetFriendUin()).Do(ctx)
			if err != nil {
				log.Error(err)
				return
			}
		}
	})

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}

}
