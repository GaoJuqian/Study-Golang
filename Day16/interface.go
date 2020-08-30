package main

import "fmt"

// 接口
type person struct {
	name string
}

func (p *person) say() {
	fmt.Printf("%v:say\n", p.name)
}

func (p *person) move() {
	fmt.Printf("%v:move\n", p.name)
}

type AAer interface { // AA类型接口
	say()
}

type BBer interface { // BB类型接口
	move()
}

type ABer interface { // 嵌套类型接口
	AAer // say()
	BBer // move()
}

func main() {

	var a AAer  // 变量a为AA接口类型
	var b BBer  // 变量b为BB接口类型
	var ab ABer // 变量ab为AB接口类型

	gjq := &person{name: "gjq"}

	a = gjq  // 变量gjq满足AA接口类型
	b = gjq  // 变量gjq满足BB接口类型
	ab = gjq // 变量gjq满足AB接口类型

	a.say()   // 执行a
	b.move()  // 执行b
	ab.move() // 执行ab嵌套
	ab.say()  // 执行ab嵌套

}
