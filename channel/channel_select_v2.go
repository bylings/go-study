package main

import (
	"fmt"
	"time"
)

func goodsC(ch chan string) {
	time.Sleep(time.Second)
	ch <- "获取商品数据接口"
}

func orderC(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "获取订单数据接口"
}

func main() {
	orderCh := make(chan string)
	goodsCh := make(chan string)
	go goodsC(goodsCh)
	go orderC(orderCh)
	// ------ 》》》 start  select  会随机选择case 中的一个执行  -----》》》
	//	select { // 检查，选择可以运行的通道进行执行（非阻塞）
	//	case order := <-orderCh:
	//		fmt.Println("order：", order)
	//	case goods := <-goodsCh:
	//		fmt.Println("goods：", goods)
	//	case <-time.After(2e9): // 设置2秒超时，常用与接口超时检测
	//		fmt.Println("已超时")
	//		//default:
	//		//fmt.Println("default")
	//	}
	// ------ 》》》 end  select  会随机选择case 中的一个执行  -----》》》

	ticker := time.After(2e9)
FCAK:
	for {
		select { // 检查，选择可以运行的通道进行执行（非阻塞）
		case order := <-orderCh:
			fmt.Println("order：", order)
			break FCAK
		case goods := <-goodsCh:
			fmt.Println("goods：", goods)
		case <-ticker: // 设置2秒超时，常用与接口超时检测
			fmt.Println("已超时")
			break FCAK
			//default:
			//fmt.Println("default")
		}
	}
	fmt.Println("8888")
}
