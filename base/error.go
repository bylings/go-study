package main

import (
	"errors"
	"fmt"
)

var errorByZero = errors.New("this value by zero")

// 相除计算
func DivisionCalculation(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errorByZero
	}
	return dividend / divisor, nil
}

// 定义结构体
type ParseError struct {
	Filename string
	Line     int
}

// 实现一个error接口，可返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一个接口
func newParseError(filename string, line int) error {
	return &ParseError{filename, line}
}

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)
	result, err := DivisionCalculation(1, 1)
	fmt.Println(result)

	// 调用error
	var e error
	e = newParseError("error.go", 1)
	fmt.Println(e.Error())

	// 
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("filename：%s  line： %d \n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")

	}
}
