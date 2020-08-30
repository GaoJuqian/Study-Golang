package main

import "fmt"

type person struct {
	name, sex string
	age       int
}

func main() {
	// 指针复习
	var a = 1
	var b = &a
	a = 8
	*b = 10
	fmt.Println(a, *b) // 10 10

	//（接受者）方法
	person1 := NewPerson("g", "男", 18)
	fmt.Println(person1)
	person1.EditAge(21)
	fmt.Println(person1)
}

// 结构体构造函数
func NewPerson(name, sex string, age int) *person {
	return &person{
		name: name,
		sex:  sex,
		age:  age,
	}
}

// 修改结构体中的age
func (p *person) EditAge(age int) {
	p.age = age
}
