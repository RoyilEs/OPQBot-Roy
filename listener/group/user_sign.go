package group

import (
	"context"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"gorm.io/gorm"
	"obqbot/global"
	"obqbot/models"
	"obqbot/models/ctype"
	"strconv"
	"time"
)

func UserSign(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		var userSign models.UserSign
		groupMsg := event.ParseGroupMsg()
		if groupMsg.ParseTextMsg().GetTextContent() == "签到" {
			//查询用户是否存在
			if groupUserThere(event, &userSign) {
				global.DB.Model(&userSign).Update("sign_type", ctype.SignOk)
				global.DB.Model(&userSign).Update("point", gorm.Expr("point + ?", 10))
				global.DB.Where("uin = ?", groupMsg.GetSenderUin()).Find(&userSign)
				time.Sleep(1 * time.Second)
				// 签到成功
				global.DB.Model(&userSign).Take(&userSign)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).
					TextMsg(
						"签到成功(￣▽￣)" +
							"\n用户: " + groupMsg.GetSenderNick() +
							"\n积分:" + strconv.FormatInt(userSign.Point, 10)).Do(ctx)
			} else {
				// 查询此用户是否签到
				err := global.DB.Take(&userSign, "sign_type = ?", ctype.SignOk).Error
				if err == nil {
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
						GroupMsg().ToUin(groupMsg.GetGroupUin()).
						TextMsg("今日已签到, 不许再签啦！").Do(ctx)
					return
				} else {
					global.DB.Model(&userSign).Update("sign_type", ctype.SignOk)
					global.DB.Model(&userSign).Update("point", gorm.Expr("point + ?", 10))
					global.DB.Where("uin = ?", groupMsg.GetSenderUin()).Find(&userSign)
					time.Sleep(1 * time.Second)
					// 签到成功
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
						GroupMsg().ToUin(groupMsg.GetGroupUin()).
						TextMsg(
							"签到成功(￣▽￣)" +
								"\n用户: " + groupMsg.GetSenderNick() +
								"\n积分:" + strconv.FormatInt(userSign.Point, 10)).Do(ctx)
				}
			}
		}
	}
}

func UserSignPoint(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		var userSign models.UserSign
		groupMsg := event.ParseGroupMsg()
		if groupMsg.ParseTextMsg().GetTextContent() == "积分" {
			// 查询此用户是否存在
			uin := groupMsg.GetSenderUin()   // 获取发送者uin
			groupUserThere(event, &userSign) //查询用户是否存在
			// 查询此用户积分
			err := global.DB.Take(&userSign, "uin = ?", uin).Error
			if err == nil {
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("目前积分:" + strconv.FormatInt(userSign.Point, 10)).Do(ctx)
			}
		}
	}
}

// 查询当前命令用户是否存在
func groupUserThere(event events.IEvent, userSign *models.UserSign) bool {
	groupMsg := event.ParseGroupMsg()
	// 查询此用户是否存在
	uin := groupMsg.GetSenderUin() // 获取发送者uin
	err := global.DB.Take(&userSign, "uin = ?", uin).Error
	if err != nil {
		// 不存在 用户入库
		sign := models.UserSign{
			NickName: groupMsg.GetSenderNick(),
			Uin:      groupMsg.GetSenderUin(),
			SignType: 0,
			Point:    0,
		}
		global.DB.Create(&sign)
		return true
	}
	return false
}
