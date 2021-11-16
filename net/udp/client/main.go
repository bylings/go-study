package main

import (
	"fmt"
	"net"
)

func main() {
	socket, _ := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 5555,
	})
	defer socket.Close()
	socket.Write([]byte("hello udp server"))

	// 接收数据
	var data [1024]byte
	n, addr, _ := socket.ReadFromUDP(data[:])
	fmt.Println("接收到的服务端数据：", string(data[:n]), addr)
}
