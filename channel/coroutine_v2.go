package main

import (
	"fmt"
)

/*
有两个协程方法，
	方法send可以循环发送指定范围的数字
	方法recv可以接收并打印出send新的；并main方法必须等两个协程完成后再结束
*/

// 阻塞之后是否影响到main执行
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sendChan(1, 10, ch)
	go recvChan(ch, done)
	<-done
	fmt.Println("888888888888")

}

func sendChan(start, end int, ch chan<- int) {
	for i := start; i < end; i++ {
		ch <- i
	}
	close(ch) // 执行完之后关闭通道
}

func recvChan(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println("num：", num)
	}
	done <- true
}
