package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(time.Now())

	wg.Add(1) // 登记协程
	go redis()
	wg.Add(2) // 登记协程(后入先出)
	go mysql()
	go file()

	// go 主程序也是一个协程 coroutine
	// 通过go 声明的方法也是协程

	wg.Wait() // 等待记录的协程执行完再执行后续的程序
	fmt.Println(time.Now())
	//time.Sleep(time.Second * 4)
}

func redis() {
	defer wg.Done() // 结束登记
	fmt.Println("start 读取redis")
	//WriteFile("/Users/mzj/Documents/code/server/go/src/test/t.txt", "写入的redis")

	//time.Sleep(time.Second * 1)
	fmt.Println("end 读取redis")
}

func mysql() {
	defer wg.Done() // 结束登记
	fmt.Println("start 读取mysql")
	//time.Sleep(time.Second * 2)
	fmt.Println("end 读取mysql")
}

func file() {
	defer wg.Done()
	fmt.Println("start 读取file")
	//time.Sleep(time.Second * 3)
	fmt.Println("end 读取file")
}

// 写入文件
func WriteFile(path, data string) (bool, error) {
	// 2、打开文件资源,设置可写并创建
	outputFile, outputError := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		return false, outputError
	}
	// 释放文件资源
	defer outputFile.Close()
	// 创建写缓冲区
	outputWriter := bufio.NewWriter(outputFile)

	// 写入文件
	_, err := outputWriter.WriteString(data)
	if err != nil {
		fmt.Println("文件写入失败", err)
		return false, nil
	}
	// 刷新缓冲区
	errs := outputWriter.Flush()
	if errs != nil {
		fmt.Println("文件刷新失败", errs)
		return false, errs
	}
	return true, nil
}
