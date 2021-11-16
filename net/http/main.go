package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle()

	fmt.Println("启动http://127.0.0.1:8888")
	// http://127.0.0.1:8888/bar
	http.HandleFunc("/bar", footHandler)

	// 监听
	//http.ListenAndServeTLS("") 监听https
	http.ListenAndServe(":8888", nil) // 监听http
}

//
func footHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("处理逻辑")
	fmt.Println("得到连接：", r.RemoteAddr)
	fmt.Println("url：", r.URL.Path)
	fmt.Println("method：", r.Method)
	w.Write([]byte("hello world"))
}
