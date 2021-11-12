package main

import (
	"fmt"
	"time"
)

func main() {
	orderCh := make(chan string)
	goodsCh := make(chan string)
	go orderChan(orderCh)
	go goodsChan(goodsCh)
REDISDONE:
	for {
		select { // 检查，选择可以运行的通道进行执行（非阻塞）
		case order := <-orderCh:
			fmt.Println("order：", order)
			break REDISDONE
		case goods := <-goodsCh:
			fmt.Println("goods：", goods)
		case <-time.After(2e9): // 设置2秒超时，常用与接口超时检测
			fmt.Println("已超时")
			break REDISDONE
			//default:
			//fmt.Println("default")
		}
	}
	fmt.Println("8888")
}

func orderChan(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "获取订单数据接口"
}

func goodsChan(ch chan string) {
	time.Sleep(time.Second)
	ch <- "获取商品数据接口"
}
