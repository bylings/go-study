package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func createImage() *image.Gray {
	// 图像大小
	const size = 300

	// 定大小创建灰度图
	var pic = image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	return pic
}

func main() {
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}

	// 使用png格式将数据写入文件
	png.Encode(file, createImage())
	file.Close()
}
