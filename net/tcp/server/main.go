package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动服务端：tcp://127.0.0.1:8888")
	// 1、监听端口，tcp://0.0.0.0:8888，监听的网络主要以本机可用ip为主
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("err：", err)
		return // return 表示程序结束
	}
	for  {
		// 2、接收客户端向服务端建立的连接
		conn, err := listen.Accept() // 可用与客户端建立连接，如果没有连接，则处于挂起阻塞状态
		if err != nil {
			fmt.Println("err :", err)
			return
		}
		// 3、处理用户的连接信息
		go handler(conn)
	}
}

func handler(c net.Conn) {
	defer c.Close() // 关闭连接
	for {
		// 获取数据
		var data [1024]byte
		n, err := bufio.NewReader(c).Read(data[:])
		if err != nil {
			fmt.Println("err：", err)
			break
		}
		fmt.Println("接收到的客户端数据：", string(data[:n]))
		c.Write([]byte("hello world ,this is server"))
	}

}
