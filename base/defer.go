package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	valueByKey      = make(map[string]int) // 定义一个演示映射
	valueByKeyGuard sync.Mutex             // 定义互斥锁
)

func readValue(key string) int {
	// 加锁
	valueByKeyGuard.Lock()

	// 读
	value := valueByKey[key]

	// 解锁
	valueByKeyGuard.Unlock()
	return value
}

func readValueV2(key string) int {
	// 加锁
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock() // 延迟解锁
	return valueByKey[key]         // 读取数据
}

/**
根据文件获取文件大小，单位/字节
*/

func getFileSize(filename string) int64 {
	// 1、打开文件
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}

	// 添加一个延迟执行释放文件句柄
	defer f.Close()
	fi, err := f.Stat() // 获得文件的元数据信息
	if err != nil {
		return 0
	}
	return fi.Size() // 单位，字节
}

func main() {
	size := getFileSize("img.png")
	fmt.Println(size)
}
