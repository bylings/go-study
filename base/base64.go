package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 需要处理的字符串
	messge := "Away from keyword. https://goland.org/"

	// 编码信息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(messge))

	// 输出编码完成的消息
	fmt.Printf(encodedMessage + "\n")

	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
