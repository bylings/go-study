package main

// 定时请求接口
import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1e9)
	for {
		<-ticker
		go orderApi()
	}
}

func orderApi() {
	fmt.Println("请求订单接口")
}
