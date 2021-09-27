package main

import (
	"fmt"
	"log"
	"sync"
)

/**
sync.Map映射，不能使用make创建
*/

func main() {
	var scene sync.Map

	scene.Store("redis", "127.0.0.1") // 使用Store创建sync.Map 映射
	scene.Store("mysql", "192.168.1.1")
	scene.Store("swoole", "192.168.1.2")

	scene.Delete("mysql")

	// 使用Load去取映射指定key
	redisValue, err := scene.Load("redis")
	if err != true {
		log.Fatal(err)
	}
	fmt.Println(redisValue)

	// 使用Range遍历sync.Map
	scene.Range(func(key, value interface{}) bool {
		fmt.Println("interface：", key, value)
		return true
	})

}
