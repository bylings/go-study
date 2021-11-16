package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("启动udp://127.0.0.1:5555")
	// 监听地址：传递net.UDPAddr对象
	listen, _ := net.ListenUDP("udp",
		&net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 5555,
		},
	)
	defer listen.Close()
	for {
		// 接收信息
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		fmt.Println("接收到的客户端数据：", string(data[:n]), addr)

		if err != nil {
			continue
		}
		// 发送信息
		listen.WriteToUDP([]byte("this is udp server"), addr)
	}
}
