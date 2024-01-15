package main

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/opq-osc/OPQBot/v2"
	"github.com/opq-osc/OPQBot/v2/events"
	core2 "obqbot/core"
	"obqbot/global"
	"obqbot/listener"
	"obqbot/listener/friend"
	"obqbot/listener/group"
	"time"
)

type WSClient struct {
	conn         *websocket.Conn
	reconnect    bool
	reconnectCh  chan struct{}
	reconnectDur time.Duration
}

func main() {
	// 读取配置文件
	core2.InitConf()
	// 初始化日志
	global.Log = core2.InitLogger()

	core, err := OPQBot.NewCore(global.OBQBotUrl)
	if err != nil {
		panic(err)
	}

	core.On(events.EventNameGroupMsg, listener.ListenGroup)

	core.On(events.EventNameGroupMsg, group.Hello)
	core.On(events.EventNameGroupMsg, group.Yian)
	core.On(events.EventNameGroupMsg, group.GoodNight)
	core.On(events.EventNameGroupMsg, group.Img)

	core.On(events.EventNameFriendMsg, friend.Hello)
	core.On(events.EventNameFriendMsg, listener.ListenFriend)

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
	// 让程序保持运行，以便能够处理WebSocket连接
	select {}
}
