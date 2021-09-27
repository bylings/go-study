package main

import (
	"image"
	"image/color"
)

func main() {
	// 图像大小
	const size = 300

	// 定大小创建灰度图
	var pic = image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}
}
