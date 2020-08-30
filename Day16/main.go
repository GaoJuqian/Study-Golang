package main

import "fmt"

// interface{} 接口
func main() {
	var a interface{} // interface{} 不知道变量是哪种数据类型

	a = "BHJJNKM"
	v, ok := a.(string)
	if ok {
		fmt.Println(v)
	}

	switch a.(type) { // 使用断言判断type
	case int:
		fmt.Printf("a的类型是int:%T\n", a)
	case string:
		fmt.Printf("a的类型是string:%T\n", a)
	case bool:
		fmt.Printf("a的类型是bool:%T\n", a)
	}

	test1("gvjhbkjnk")

	// 使用接口创建map 任意类型value
	test2 := make(map[string]interface{}, 5)
	fmt.Println(test2)
}

func test1(test interface{}) { // 用于函数
	switch test.(type) {
	case int:
		fmt.Printf("test1的类型是int:%T\n", test)
	case string:
		fmt.Printf("test1的类型是string:%T\n", test)
	case bool:
		fmt.Printf("test1的类型是bool:%T\n", test)
	}
}
