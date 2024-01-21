package group

import (
	"context"
	"encoding/base64"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"io/ioutil"
	"obqbot/global"
	"obqbot/utils"
)

const MyGOPath = "uploads/mygo/"

var MyGO = map[int]string{
	0:  "処救生",
	1:  "名無声",
	2:  "壱雫空",
	3:  "影色舞",
	4:  "春日影",
	5:  "栞",
	6:  "歌いましょう鳴らしましょう",
	7:  "無路矢",
	8:  "碧天伴走",
	9:  "詩超絆",
	10: "迷星叫",
	11: "迷路日々",
}

func randGetMyGO() string {
	// 获取map的长度
	length := len(MyGO)
	randomIndex := utils.Random(0, length)
	// 使用随机索引从map中获取元素
	randomElement, ok := MyGO[randomIndex]
	if ok {
		return randomElement
	} else {
		return ""
	}
}

func RandMusicMyGo(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		// 获取群聊体
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ParseTextMsg().GetTextContent()
		if text == "MyGo" {
			myGO := randGetMyGO()
			file, err := ioutil.ReadFile(MyGOPath + myGO + ".mp3")
			if err != nil {
				global.Log.Error(err)
				return
			}
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg(myGO + ":[]正在下载中。。。 耐心等待").Do(ctx)
			voice, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().
				GroupVoice().SetBase64Buf(base64.StdEncoding.EncodeToString(file)).DoUpload(ctx)
			if err != nil {
				global.Log.Error(err)
				return
			}
			log.Info("播放音乐：" + myGO)
			log.Debug(voice)
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
				GroupMsg().ToUin(groupMsg.GetGroupUin()).VoiceMsg(voice).Do(ctx)
		}
	}
}
