package main

import (
	"fmt"
)

// 变量、常量

func main() {
	//	变量声明
	var name string
	var age int
	var isOK bool
	fmt.Println(name, age, isOK)
	//	统一变量声明
	var (
		name1 string
		age1  int
		isOK1 bool
	)
	fmt.Println(name1, age1, isOK1)
	//	编译器类型推倒
	var name3 = "naive"
	var age3 = 21
	var isOK3 = false
	fmt.Printf("类型:%T,%T,%T \n", name3, age3, isOK3)
	//	短变量声明(仅在函数内生效)
	name4 := "naive"
	age4 := 21
	fmt.Println(name4, age4)

	/* @run
	 0 false
	 0 false
	类型:string,int,bool
	naive 21
	*/
	//	匿名变量
	// _ 短下划线，用于忽略值、站位，不占用内存，命名空间
	fmt.Println(a, b, c) // 输出相同
	idea := 1301832
	fmt.Printf("%T\n", idea)
}

// 全局定义
const (
	a = "gaojuain"
	b
	c
)
