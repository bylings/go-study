package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
filename			ini文件名
expectSection		期望读取的字段
expectKey			期望读取字段中的键
*/
func getValue(filename, expectSection, expectKey string) string {
	// 1、读取文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	// 在函数结束是，关闭文件
	defer file.Close()

	// 读取行文本
	reader := bufio.NewReader(file)

	// 当前读取的段的名字
	var sectionName string
	for {
		// 读取文件的一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// 切掉左右两边的空白符
		linestr = strings.TrimSpace(linestr)

		// 忽略空行
		if linestr == "" {
			continue
		}

		// 忽略注释
		if linestr[0:] == ";" {
			continue
		}

		// 读取段和键值的代码

		linestrLen := len(linestr)
		if linestr[0] == '[' && linestr[linestrLen-1:] == "]" {
			// 将段名取出
			sectionName = linestr[1 : linestrLen-1]

			// 读取键值

		} else if sectionName == expectSection {
			// 切开等号分割的键值对
			pair := strings.Split(linestr, "=")

			// 保证切开只有1个等号分割的键值情况
			if len(pair) == 2 {
				// 去除多余的空白字符
				key := strings.TrimSpace(pair[0])

				// 判断是期望的值
				if key == expectKey {
					return strings.TrimSpace(pair[1])
				}
			}

		}
	}
	return ""
}

// 从ini配置文件中查询需要的值
func main() {
	inivalue := getValue("php.ini", "redis", "extension")
	fmt.Println(inivalue)
}
