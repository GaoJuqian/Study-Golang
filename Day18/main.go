package main

import (
	"fmt"
	"reflect"
)

// 反射
func main() {
	var a int32 = 16
	reflectType(&a) // 需要修改可以传指针
	fmt.Println(a)

	var s []int
	s = append(s, 1)
	fmt.Println(s)

}

func reflectType(v interface{}) {
	test1 := reflect.TypeOf(v)  // reflect.TypeOf 类型
	test2 := reflect.ValueOf(v) // reflect.ValueOf 值
	test2.Elem().SetInt(100)    // Elem()操作指针.SetInt()设置值
	fmt.Println(test1)
	fmt.Println(test2)
}
