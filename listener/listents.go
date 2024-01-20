package listener

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	"obqbot/utils"
	"strconv"
)

func ListenGroup(_ context.Context, event events.IEvent) {
	groupMsg := event.ParseGroupMsg() // 群消息
	s := "群:" + groupMsg.GetGroupInfo().GroupName
	if !utils.IsInGroupS(groupMsg.GetGroupUin(), global.NotListentsGroupUids) {
		groupMsgNickUin := "成员:" + groupMsg.GetSenderNick() + "(" + strconv.FormatInt(groupMsg.GetSenderUin(), 10) + ")"
		if !groupMsg.IsFromBot() {
			global.Log.Info(s + "[" + strconv.FormatInt(groupMsg.GetGroupUin(), 10) + "]")
			global.Log.Info(groupMsgNickUin)
			uin := groupMsg.GetGroupUin()
			textGroupContent := groupMsg.ParseTextMsg().GetTextContent()
			if textGroupContent == "" {
				textGroupContent = "[图片]"
			}
			log.Info("群信息:" + textGroupContent + "(" + strconv.FormatInt(uin, 10) + ")")
		}
	}
}

func ListenFriend(_ context.Context, event events.IEvent) {
	friendMsg := event.ParseFriendMsg() // 好友消息
	friendMsgNickUin := "好友:" + friendMsg.GetFriendUid() + "(" + strconv.FormatInt(friendMsg.GetSenderUin(), 10) + ")"
	global.Log.Info(friendMsgNickUin)
	friendUin := friendMsg.GetFriendUin()
	textFriendContent := friendMsg.ParseTextMsg().GetTextContent()
	if textFriendContent == "" {
		textFriendContent = "[图片]"
	}
	log.Info("好友信息:" + textFriendContent + "(" + strconv.FormatInt(friendUin, 10) + ")")

}
