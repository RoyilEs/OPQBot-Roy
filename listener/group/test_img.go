package group

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"obqbot/global"
	"time"
)

func Img(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ExcludeAtInfo().ParseTextMsg().GetTextContent()
		if groupMsg.AtBot() {
			if text == "test" {
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).TextMsg("test").DoWithCallBack(ctx, func(iApiBuilder *apiBuilder.Response, err error) {
					response, err := iApiBuilder.GetGroupMessageResponse()
					if err != nil {
						return
					}
					time.Sleep(time.Second * 1)
					apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).GroupManager().RevokeMsg().ToGUin(groupMsg.GetGroupUin()).MsgSeq(response.MsgSeq).MsgRandom(response.MsgTime).Do(ctx)
				})
			}
			if text == "music" {
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).JsonMsg("{\"app\":\"com.tencent.structmsg\",\"config\":{\"ctime\":1705645701,\"forward\":1,\"token\":\"84bf7ed2c28088380be7b0d26965b816\",\"type\":\"normal\"},\"extra\":{\"app_type\":1,\"appid\":100495085,\"msg_seq\":7325692505905919018,\"uin\":3092179918},\"meta\":{\"music\":{\"action\":\"\",\"android_pkg_name\":\"\",\"app_type\":1,\"appid\":100495085,\"ctime\":1705645701,\"desc\":\"黄龄/HOYO-MiX\",\"jumpUrl\":\"https://y.music.163.com/m/song?id=1971144922&uct2=MqhtmIGvKeEwpsKh2%2BlxyA%3D%3D&dlt=0846&app_version=9.0.10\",\"musicUrl\":\"http://music.163.com/song/media/outer/url?id=1971144922&userid=1469868715&sc=wm&tn=\",\"preview\":\"http://p2.music.126.net/cu9T_JCh5mt3aipWJoy03w==/109951167767293721.jpg?imageView=1&thumbnail=1440z3040&type=webp&quality=80\",\"sourceMsgId\":\"0\",\"source_icon\":\"https://i.gtimg.cn/open/app_icon/00/49/50/85/100495085_100_m.png\",\"source_url\":\"\",\"tag\":\"网易云音乐\",\"title\":\"TruE (崩坏3《因你而在的故事》动画短片印象曲)\",\"uin\":3092179918}},\"prompt\":\"[分享]TruE (崩坏3《因你而在的故事》动画短片印象曲)\",\"ver\":\"0.0.0.1\",\"view\":\"music\"}").Do(ctx)
				apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().
					GroupMsg().ToUin(groupMsg.GetGroupUin()).JsonMsg("{\"app\":\"com.tencent.creategroupmsg\",\"config\":{\"autosize\":true,\"ctime\":1679744821,\"forward\":true,\"token\":\"056dccd5079a4fed285efa3c313aa260\",\"type\":\"normal\"},\"desc\":\"掌上助手\",\"key\":\"058ef0f2ada7db3576bd0a1ade4299a9\",\"meta\":{\"groupinfo\":{\"cateid\":23,\"desc\":\"★★掌上助手★★\\r\\n插件设置  发送模式  克隆管理\\r\\n主人问答  原神语音  文件中转\\r\\n通知系统  我的设定  自动授权\\r\\n卡片管理  群聊互动  主人功能\\r\\n自助头衔  更新插件  静默模式\\r\\n自动邮件  授权管理  窥屏检测\\r\\n自动备注  定时上线  音乐系统\\r\\n艾特回复  自身撤回  伪造聊天\\r\\n群发好友  群发通知  续火功能\\r\\n定时群发  特色群发  运行时长\\r\\n空间秒赞  空间评论  更新插件\\r\\n-----------\\r\\nPs：\\r\\n发送 '运行时长'\",\"learn_mode\":0,\"status\":0,\"subid\":10011,\"troopnum\":\"566725583\"}},\"prompt\":\"您已被移除群聊\",\"ver\":\"0.0.0.1\",\"view\":\"main\"}").Do(ctx)
			}
		}
		if text == "img" {
			//client := Ok3Http.NewHTTPClient("https://i.pixiv.re/img-original/img/2020/08/27/03/29/53/83959163_p1.jpg")
			////body, err := client.DoGet("", nil)
			//if err != nil {
			//	panic(err)
			//}
			pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().GroupPic().SetFileUrlPath("https://i.pixiv.re/img-original/img/2024/01/12/03/01/42/115092758_p0.jpg").DoUpload(ctx)
			if err != nil {
				panic(err)
			}
			log.Debug(pic)
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).Do(ctx)
		}
	}

}
