package main

import "fmt"

// 类型注释
type Man struct {
	name   string
	length int64
}

type people struct {
	name, sex string
	age       int
}

type aint = int  // 类型别名
type myint = int // 类型定义

// 自定义类型
func main() {

	var people Man
	fmt.Printf("%T\n %#v\n", people, people)

	test1 := Man{}
	fmt.Printf("%T\n %#v\n", test1, test1)

	//	指针结构体
	var people1 = new(Man)
	people1.name = "a"
	fmt.Printf("%T\n %#v\n", people1, people1)

	test2 := &Man{}
	test2.name = "b"
	fmt.Printf("%T\n %#v\n", test2, test2)

	//	结构体构造函数
	a := newPeople("gg", "nan", 18)
	fmt.Printf("%T\n %v\n", a, a)
}

func newPeople(name, sex string, age int) *people {
	return &people{
		name: name,
		sex:  sex,
		age:  age,
	}
}
