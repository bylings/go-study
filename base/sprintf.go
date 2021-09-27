package main

import (
	"fmt"
	"math"
)

func main() {

	// 两个参数格式化
	progress := 2
	targer := 8

	title := fmt.Sprintf("已收集%d个药材，还需要完成%d个任务", progress, targer)
	fmt.Println(title)

	// 按照数值本身的格式输出
	pi:=math.Pi
	variant:=fmt.Sprintf("%v  %v  %v","张三李四",pi,true)
	fmt.Println(variant)

	// 匿名结构体声明
	profile:=&struct {
		Name string
		age int
	}{
		Name: "张三",
		age: 11,
	}

	/**
	%v 		按照本来的值输出
	%+v 	在%v基础上，对结构体字段名和值进行展开
	%#v		输出Go语言语法格式的值
	%T		输出Go语言语法格式的类型和值
	*/
	fmt.Printf("使用'%%v 输出本来的值'  %v\n",profile)
	fmt.Printf("使用'%%+v 对结构体的字段和值进行展开'  %+v\n",profile)
	fmt.Printf("使用'%%#v 输出语法格式的值'  %#v\n",profile)
	fmt.Printf("使用 '%%T 语法格式的类型'  %T\n",profile)
}
