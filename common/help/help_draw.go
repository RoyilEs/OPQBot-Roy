package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	width, height := 200, 200
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// 绘制红色边框
	for x := 0; x < width; x++ {
		img.Set(x, 0, color.RGBA{255, 0, 0, 255})
		img.Set(x, height-1, color.RGBA{255, 0, 0, 255})
	}
	for y := 0; y < height; y++ {
		img.Set(0, y, color.RGBA{255, 0, 0, 255})
		img.Set(width-1, y, color.RGBA{255, 0, 0, 255})
	}

	// 在图像中央写入文字
	text := "/help 帮助" +
		"\n\t\t\t\t\t来点粥图 来点舟图(两种模式 速度与画质不同)" +
		"\n\t\t\t\t\t晚安" +
		"\n\t\t\t\t\t一言" +
		"\n\t\t\t\t\t" +
		"cnm \"tag\" 根据tag指定涩图(暂只支持但tag)"

	col := color.Black
	point := fixed.Point26_6{
		X: fixed.Int26_6(500),
		Y: fixed.Int26_6(1000),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)

	// 将图像保存为PNG文件
	outFile, _ := os.Create("red_square_with_text.png")
	defer outFile.Close()
	png.Encode(outFile, img)
}
