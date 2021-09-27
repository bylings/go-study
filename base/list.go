package main

import (
	"container/list"
	"fmt"
)

/**
列表
*/
func main() {
	/**
	初始化 list
	1、var newList=list.New()
	2、newList:=list.list
	*/
	newList := list.New()
	newList.PushFront("a")           // 从列表前方插入元素
	newList.PushBack("b")            // 从列表后方插入元素
	element := newList.PushBack("c") // 从列表后方插入元素

	newList.InsertBefore("d", element) // 从指定句柄之前插入元素
	newList.InsertAfter("e", element)  // 从指定句柄插入元素

	// 删除指定句柄
	newList.Remove(element)

	// 遍历列表，当遇到nil时退出
	for i := newList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
