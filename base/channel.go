package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		 创建通道
		通道变量名，变量声明的是建议具有 ch 或 chan
	*/
	ch := make(chan string) // 创建通道  - 字符串类型通道
	go send(ch)
	go recv(ch)
	time.Sleep(time.Second * 2)
}

func send(ch chan string) { // 通道参数引用
	ch <- "aaaaa"
	fmt.Println("111")
	ch <- "bbbbb"
	fmt.Println("222")
	ch <- "ccccc"
	fmt.Println("333")
}

func recv(ch chan string) { // 通道参数引用
	ach := <-ch
	fmt.Println("ach：", ach)

	fmt.Println("接收到的通道信息：", <-ch)
	//fmt.Println("接收到的通道信息：", <-ch)

	if <-ch == "ccccc" {
		fmt.Println("接收到的：ccccc")
	}
}
