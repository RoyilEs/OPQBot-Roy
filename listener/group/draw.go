package group

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"github.com/opq-osc/OPQBot/v2/events"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"obqbot/global"
)

func drawAndEncodeTobase64(test string) (string, error) {
	log.Info("1.开始绘制图片")
	width := 80
	height := 60
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	background := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	draw.Draw(m, m.Bounds(), &image.Uniform{C: background}, image.ZP, draw.Src)

	// 添加文本到图像
	d := &font.Drawer{
		Dst:  m,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13, // 假设我们有一个可用的基本字体
	}
	d.Dot = fixed.Point26_6{X: fixed.Int26_6(width / 2), Y: fixed.Int26_6(height / 2)} // 文本中心位置
	d.DrawString(test)

	var buf bytes.Buffer
	err := jpeg.Encode(&buf, m, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encode JPEG: %v", err)
	}
	encodeToString := base64.StdEncoding.EncodeToString(buf.Bytes())
	log.Info("2.绘制图片完成")
	return encodeToString, nil
}

func Draw(ctx context.Context, event events.IEvent) {
	if event.GetMsgType() == events.MsgTypeGroupMsg {
		groupMsg := event.ParseGroupMsg()
		text := groupMsg.ExcludeAtInfo().ParseTextMsg().GetTextContent()
		if text == "draw" {
			tobase64, err := drawAndEncodeTobase64("Hello, World!")
			if err != nil {
				global.Log.Error(err)
				return
			}
			pic, err := apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).Upload().GroupPic().SetBase64Buf(tobase64).DoUpload(ctx)
			if err != nil {
				global.Log.Error(err)
				return
			}
			apiBuilder.New(global.OBQBotUrl, event.GetCurrentQQ()).SendMsg().GroupMsg().ToUin(groupMsg.GetGroupUin()).PicMsg(pic).Do(ctx)
		}
	}
}
