package main

import "fmt"

/**
复制切片
*/

func main() {
	// 设置元素数量
	const elementCount = 1000

	// 预分配元素足够多的元素切片
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = 1
	}
	// 引用切片数据（这里同时引用的是地址）
	refData := srcData

	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)

	// 复制到新的切片中
	copy(copyData, srcData)

	// 修改原始数据的第一个值
	srcData[0] = 999
	srcData[2] = 666

	fmt.Println(refData[0])

	// 打印第一个和最后一个复制的切片，由于切片下标是从0开始，所以最有一个下标就是len(slice)-1
	fmt.Println(copyData[0], copyData[len(copyData)-1])

	// 复制原始数据从4到6（不包含6）
	copy(copyData, srcData[4:6])

	fmt.Println(len(copyData))

	for i := 0; i < 5; i++ {
		fmt.Println(copyData[i])
	}
	fmt.Println(refData[2])
}
