package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/opq-osc/OPQBot/v2/apiBuilder"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"obqbot/global"
	"obqbot/models/pixiv"
	"os"
	"strconv"
)

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

func drawTextOnImage(img image.Image, text string, fontPath string, fontSize float64, color color.Color) (image.Image, error) {
	face, err := loadFontFace(fontPath, fontSize)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	d := &font.Drawer{
		Dst:  newImg,
		Src:  image.NewUniform(color),
		Face: face,
	}

	// 计算文本绘制位置 - 右下角
	textD := d.MeasureString(text).Round()
	x := bounds.Dx() - textD
	y := bounds.Dy() - 10

	pt := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	d.Dot = pt
	d.DrawString(text)

	return newImg, nil
}

func downloadImage(url string) (image.Image, int, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, 0, err
	}
	defer resp.Body.Close()

	imgData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, 0, err
	}
	height, width := apiBuilder.GetPicHW(imgData)

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, 0, 0, err
	}

	return img, height, width, nil
}

// drawAndEncodeToBase64 空白画图demo
func drawAndEncodeToBase64(text string) (string, error) {
	// 创建一个空白图像
	width := 800
	height := 600
	m := image.NewRGBA(image.Rect(0, 0, width, height))
	background := color.RGBA{255, 255, 255, 255} // 白色背景
	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	// 添加文本到图像
	d := &font.Drawer{
		Dst:  m,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13, // 假设我们有一个可用的基本字体
	}
	d.Dot = fixed.Point26_6{X: fixed.Int26_6(width / 2), Y: fixed.Int26_6(height / 2)} // 文本中心位置
	d.DrawString(text)                                                                 // 绘制文本

	var buf bytes.Buffer
	err := jpeg.Encode(&buf, m, nil)
	if err != nil {
		return "", fmt.Errorf("failed to encode JPEG: %v", err)
	}

	base64Img := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Img, nil
}

func main() {
	//base64Str, err := drawAndEncodeToBase64("Hello, World!")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Base64 encoded image: %s\n", base64Str)
	query, _ := pixiv.NewPixiv().Set().DoQuery()
	pixivResponse, _ := pixiv.NewPixiv().Do(pixiv.PixivUrl, query)
	data := pixivResponse.GetData()[0]
	size := data.GetDataUrls().GetSize()
	fmt.Println(size)
	i, _, _, err := downloadImage(size)
	imageFilePath := "uploads/test.jpg"
	fontFilePath := "uploads/Font.ttf"
	text := "Pixiv|" + strconv.FormatInt(data.GetDataUid(), 10) + "|" + data.GetDataAuthor()
	color1 := color.RGBA{
		R: 255,
		G: 0,
		B: 100,
		A: 240,
	}
	//x := 400
	//y := 1000
	// 加载图片
	_, err = loadImage(imageFilePath)
	if err != nil {
		panic(err)
	}
	newImg, err := drawTextOnImage(i, text, fontFilePath, 50, color1)
	if err != nil {
		global.Log.Error(err)
		return
	}
	// 现在newImg包含了带有文本的新图片，你可以选择保存它
	outFile, err := os.Create("output3.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// 编码并写入JPEG格式
	err = jpeg.Encode(outFile, newImg, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully drew text on the image and saved as output.jpg")
}
