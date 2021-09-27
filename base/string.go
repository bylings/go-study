package main

import (
	"bytes"
	"fmt"
)

func stringBuilder() {
	hammer := "张三"
	sickle := "李四"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer

	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	fmt.Println(stringBuilder.String())
}

func main() {
	//stringBuilder()

	fmt.Println("张三" + "李四")
}
