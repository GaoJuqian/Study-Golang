package main

// 结构体 嵌套 继承
import "fmt"

// test1
type User struct {
	name      string
	sex       string
	age       int
	*UserInfo // 匿名指针结构体嵌套
}

// test2
type UserInfo struct {
	phone int
}

func (u *User) SayHi() {
	fmt.Printf("Hi :%v\n", u.name)
}
func (u *UserInfo) OutPhone() {
	fmt.Printf("phone:%v\n", u.phone)
}

func main() {
	a := &User{
		name: "A",
		sex:  "man",
		age:  18,
		UserInfo: &UserInfo{
			phone: 123123123123,
		},
	}
	a.SayHi()
	// 结构体继承
	a.OutPhone()

}
