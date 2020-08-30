package main

import "fmt"

//
func main() {
	s := &Student{"aa", &Student{"bb", nil}}
	test(s)
	fmt.Println(s.Name)

}

type Student struct {
	Name string
	Next *Student
}

// 函数的接收和返回都是值拷贝
func test(s *Student) {
	//s = s.Next //值拷贝操作
	*s = *s.Next
	//s.Name = "aaa" // 值拷贝导致 s.Name(*s.Name)语法糖失效
}
