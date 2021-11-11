package main

import (
	"fmt"
	"runtime"
	"time"
)

// main  结束之后，内部所有的携程都不再运行
func main() {
	//go Run1()
	//go Run2()
	//time.Sleep(time.Second)

	// 设置当前程序并发时占用的CPU逻辑核数，默认为系统核数
	runtime.GOMAXPROCS(3)
	go C()
	go B()
	go A()

	time.Sleep(time.Second)
}

func Run1() {
	fmt.Println("start Run1")
	runtime.Gosched() // 让出执行权
	fmt.Println("end Run1")
}

func Run2() {
	fmt.Println("start Run2")
	runtime.Goexit() // 结束协程运行
	fmt.Println("end Run2")
}

func A() {
	fmt.Println("A")
}

func B() {
	fmt.Println("B")
}

func C() {
	fmt.Println("C")
}
