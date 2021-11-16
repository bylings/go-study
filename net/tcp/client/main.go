package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、创建建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err ：", err)
		return
	}
	fmt.Println("与tcp://0.0.0.0:8888建立连接")
	defer conn.Close()
	// 2、进行数据接收和发送

	conn.Write([]byte("你好 server"))
	var data [1024]byte
	n, _ := conn.Read(data[:])
	fmt.Println("接收到server的数据：", string(data[:n]))

	// 3、关闭连接 | 不关闭不影响，如果服务端没有心跳会出错
}
