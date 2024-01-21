package utils

import (
	"fmt"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
)

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func IsInGroupS(str int64, groupUids []int64) bool {
	for _, s := range groupUids {
		if s == str {
			return true
		}
	}
	return false
}

func IsInListToS(str any, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func IsAdmins(uid int64, adminUids []int64) bool {
	for _, s := range adminUids {
		if s == uid {
			return true
		}
	}
	return false
}

func ContainsURL(str string) bool {
	// 定义匹配网址的正则表达式模式
	pattern := `(?i)(https?:\/\/)?([\w-]+\.)?([\w-]+\.[\w-]+)`
	// 编译正则表达式模式
	regex := regexp.MustCompile(pattern)
	// 使用正则表达式进行匹配
	return regex.MatchString(str)
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func loadFontFace(fontPath string, fontSize float64) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	parse, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %v", err)
	}

	hinting := font.HintingFull
	face, _ := opentype.NewFace(parse, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: hinting,
	})

	return face, nil
}

func drawTextOnImage(img image.Image, text string, fontPath string, fontSize float64, color color.Color, x, y int) (image.Image, error) {
	// 加载字体
	face, err := loadFontFace(fontPath, fontSize)
	if err != nil {
		return nil, err
	}

	// 创建一个新的图像以便于绘制
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)

	// 将原始图片复制到新图像上
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	d := &font.Drawer{
		Dst:  newImg,
		Src:  image.NewUniform(color),
		Face: face,
	}

	// 设置绘制点的位置
	pt := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	// 在指定位置绘制文本
	d.Dot = pt
	d.DrawString(text)

	return newImg, nil
}
