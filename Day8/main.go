package main

import (
	"fmt"
	"strings"
)

// 匿名函数，闭包
func main() {
	// 匿名函数
	a := func() {
		fmt.Println("这是a匿名函数")
	}
	a()
	//	匿名函数
	func() {
		fmt.Println("这是直接执行的匿名函数")
	}()
	//	闭包 函数调用外层变量
	// test1
	name := test1("a")
	name()
	// test2
	FileName := test2("aaa")
	out := FileName(".img")
	fmt.Println(out)
	//	test3

}
func test1(name string) func() {
	return func() {
		fmt.Println("你的名字是：" + name)
	}
}
func test2(FileName string) func(end string) string {
	return func(end string) string {
		if !strings.HasSuffix(FileName, end) {
			return FileName + end
		}
		return FileName
	}
}
