package group

import (
	"context"
	"fmt"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"math/rand"
	"obqbot/global"
	"obqbot/models"
	"obqbot/utils"
	"regexp"
	"strings"
	"time"
)

func InitMeme() {
	friendTag := models.FriendTag{
		Name: "rtwyzz",
		TagsData: models.StringArray{
			"我操死你的吗rtwyzz",
			"怂逼,不敢对嘴",
		},
	}
	global.DB.Create(&friendTag)

	global.DB.Where("name = ?", "rwtyzz").First(&friendTag)
	fmt.Println(friendTag.TagsData[0])
}

func NewTag(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		if utils.IsAdmins(groupMsg.GetSenderUin(), global.AdminUids) {
			texts := strings.Split(text, " ")
			if texts[0] == "newMeme" && len(texts) == 2 {
				err := global.DB.Where("name = ?", texts[1]).First(&models.FriendTag{}).Error
				if err == nil {
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
						GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("此Meme已存在").Do(ctx)
					return
				}
				friendTag := models.FriendTag{
					Name: texts[1],
				}
				global.DB.Create(&friendTag)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg(`
成功添加新Meme
添加词条请使用setNameTag "词条名" "词条内容"
`).Do(ctx)
			}
		}
	}
}

func SetNameTag(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		if utils.IsAdmins(groupMsg.GetSenderUin(), global.AdminUids) {
			texts := strings.Split(text, " ")
			if len(texts) == 3 && texts[0] == "setNameTag" {
				// 新增
				var friendTag models.FriendTag
				global.DB.Where("name = ?", texts[1]).First(&friendTag)
				friendTag.TagsData = append(friendTag.TagsData, texts[2])
				global.DB.Save(&friendTag)

				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("设置成功").Do(ctx)
			}
		}
	}
}

func DeleteNameTag(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		if utils.IsAdmins(groupMsg.GetSenderUin(), global.AdminUids) {
			texts := strings.Split(text, " ")
			if len(texts) == 3 && texts[0] == "deleteNameTag" {
				// 删除
				var (
					friendTag models.FriendTag
				)
				err := global.DB.Where("name = ?", texts[1]).First(&friendTag).Error
				if err != nil {
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
						GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("此Meme不存在").Do(ctx)
					return
				} else {
					temp := false
					for _, v := range friendTag.TagsData {
						if v == texts[2] {
							temp = true
							break
						}
					}
					if !temp {
						apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
							GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("压根没这个Tag你删什么呢?").Do(ctx)
						return
					}
				}
				for i, v := range friendTag.TagsData {
					if v == texts[2] {
						friendTag.TagsData = append(friendTag.TagsData[:i], friendTag.TagsData[i+1:]...)
						break
					}
				}
				global.DB.Save(&friendTag)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("删除成功").Do(ctx)
			}
		}
	}
}

func RandomTag(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		pattern := regexp.MustCompile(`来点(.*)`)
		matches := pattern.FindStringSubmatch(text)
		if len(matches) > 1 {
			result := matches[1]
			var (
				friendTag models.FriendTag
			)
			global.DB.Where("name = ?", result).First(&friendTag)
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(friendTag.TagsData))
			tag := friendTag.TagsData[randomIndex]
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg(tag).Do(ctx)
		}
	}
}

func AllMemeTag(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		texts := strings.Split(text, " ")
		if texts[0] == "allMeme" {
			var (
				friendTag models.FriendTag
			)
			global.DB.Where("name = ?", texts[1]).First(&friendTag)
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg(strings.Join(friendTag.TagsData, "\n")).Do(ctx)
		}
	}
}
