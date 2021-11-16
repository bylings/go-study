package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func main() {
	http.Handle("/", websocket.Handler(server))
	// 设定监听
	http.ListenAndServe(":7777", nil)
}

// 当有连接进来时，底层自动触发
func server(ws *websocket.Conn) {
	fmt.Println("new connet")

	data := make([]byte, 1024)
	for {
		d, err := ws.Read(data)
		if err != nil {
			fmt.Println("err ：", err)
			break
		}
		fmt.Println("读取到的信息：", d)
		ws.Write([]byte("this is websocker server"))
	}
}
