package group

import (
	"context"
	"fmt"
	"github.com/opq-osc/OPQBot/v2"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"golang.org/x/sync/errgroup"
	"obqbot/global"
	"obqbot/utils"
	"strings"
)

var apiUrl = global.OBQBotUrl

func HandleGroupMsg(core *OPQBot.Core) {
	// help
	core.On(events.EventNameGroupMsg, func(ctx context.Context, event events.IEvent) {
		if event.GetMsgType() == events.MsgTypeGroupMsg {
			groupMsg := event.ParseGroupMsg()
			if groupMsg.ParseTextMsg().GetTextContent() == "/help" {
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg(`
	/help 帮助
	来点粥图 来点舟图(两种模式 速度与画质不同)
	晚安	
	一言
	cnm "tag" 根据tag指定涩图(暂只支持单tag)

——作者比较废物就写了这么点
										`).Do(ctx)
			}
		}
	})
	// 欢迎公告
	core.On(events.EventNameGroupJoin, func(ctx context.Context, event events.IEvent) {
		groupMsg := event.PraseGroupJoinEvent()
		apiBuilder.New(apiUrl, event.GetCurrentQQ()).SendMsg().
			GroupMsg().ToUin(groupMsg.GetGroupUId()).
			TextMsg("欢迎新人~").Do(ctx)
		apiBuilder.New(apiUrl, event.GetCurrentQQ()).GroupManager().
			GetGroupMemberLists(groupMsg.GetGroupUId(), "").Do(ctx)
	})
	//指令 禁言@用户 1h (禁言用户和时间中间需要有一个空格)
	core.On(events.EventNameGroupMsg, func(ctx context.Context, event events.IEvent) {
		if event.GetMsgType() == events.MsgTypeGroupMsg {
			groupMsg := event.ParseGroupMsg()
			if utils.IsAdmins(groupMsg.GetSenderUin(), global.AdminUids) {
				// 信息体
				text := groupMsg.ExcludeAtInfo().ParseTextMsg().GetTextContent()
				// 禁言用户
				if strings.HasPrefix(text, "禁言") && groupMsg.ContainedAt() {
					// 获取被禁言的用户
					user := groupMsg.GetAtInfo()
					if len(user) == 0 {
						apiBuilder.New(apiUrl, event.GetCurrentQQ()).SendMsg().
							GroupMsg().ToUin(groupMsg.GetGroupUin()).
							TextMsg("格式错误，@自带的空格不算哦").Do(ctx)
						return
					}
					// 获取禁言时间
					time := strings.Split(text, " ")[2]
					intTime := 0
					switch time {
					case "1m":
						intTime = 60
					case "1h":
						intTime = 3600
					case "1d":
						intTime = 86400
					case "1w":
						intTime = 604800
					}
					if len(user) > 0 {
						group, ctx := errgroup.WithContext(context.Background())
						for _, u := range user {
							group.Go(func() error {
								return apiBuilder.New(apiUrl, event.GetCurrentQQ()).GroupManager().
									ProhibitedUser().ToGUin(groupMsg.GetGroupUin()).ToUid(groupMsg.GetSenderUid()).ShutTime(intTime).Do(ctx)
							})
							group.Go(func() error {
								global.Log.Infof("用户%s已经被成功禁言 %s", u.Nick, time)
								return apiBuilder.New(apiUrl, event.GetCurrentQQ()).SendMsg().
									GroupMsg().ToUin(groupMsg.GetGroupUin()).
									TextMsg(fmt.Sprintf("用户%s已经被成功禁言 %s", u.Nick, time)).Do(ctx)
							})
						}
						if err := group.Wait(); err != nil {
							global.Log.Fatalf("禁言失败，未找到用户")
							apiBuilder.New(apiUrl, event.GetCurrentQQ()).SendMsg().
								GroupMsg().ToUin(groupMsg.GetGroupUin()).
								TextMsg("禁言失败，未找到用户").Do(ctx)
						}
					} else {
						global.Log.Fatalf("禁言失败，未找到用户")
						apiBuilder.New(apiUrl, event.GetCurrentQQ()).SendMsg().
							GroupMsg().ToUin(groupMsg.GetGroupUin()).
							TextMsg("禁言失败，未找到用户").Do(ctx)
					}
				}
			}
		}
	})
}
