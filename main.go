package main

import (
	"context"
	"github.com/opq-osc/OPQBot/v2"
	"github.com/opq-osc/OPQBot/v2/events"
	"github.com/robfig/cron"
	core2 "obqbot/core"
	"obqbot/flag"
	"obqbot/global"
	"obqbot/listener"
	"obqbot/listener/friend"
	"obqbot/listener/group"
)

func main() {
	// 读取配置文件
	core2.InitConf()
	// 初始化日志
	global.Log = core2.InitLogger()
	// 初始化数据库
	global.DB = core2.InitGorm()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	c := cron.New()

	err := c.AddFunc("0 0 0 * * *", group.ResetSignSignNo)
	if err != nil {
		panic(err)
	}
	c.Start()
	core, err := OPQBot.NewCore(global.OBQBotUrl)
	if err != nil {
		panic(err)
	}

	go group.HandleGroupMsg(core)
	core.On(events.EventNameGroupMsg, listener.ListenGroup)

	core.On(events.EventNameFriendMsg, listener.ListenFriend)

	core.On(events.EventNameGroupMsg, group.Hello)
	core.On(events.EventNameGroupMsg, group.Yian)
	core.On(events.EventNameGroupMsg, group.GoodNight)
	core.On(events.EventNameGroupMsg, group.Img)
	core.On(events.EventNameGroupMsg, group.Draw)
	core.On(events.EventNameGroupMsg, group.ArknightsImg)
	core.On(events.EventNameGroupMsg, group.PixivImg)

	core.On(events.EventNameGroupMsg, group.UserSign)
	core.On(events.EventNameGroupMsg, group.UserSignPoint)

	core.On(events.EventNameFriendMsg, friend.Hello)

	err = core.ListenAndWait(context.Background())
	if err != nil {
		panic(err)
	}
	// 让程序保持运行，以便能够处理WebSocket连接
	select {}
}
