package main

import "fmt"

// 数组 数组在声明时*固定数量 不可改变
func main() {
	var a [3]int
	fmt.Println(a) // [0 0 0]
	b := [3]int{}
	b[0] = 1
	fmt.Println(b) // [1 0 0]
	// ... 推导array长度
	c := [...]bool{true, false, false, true}
	fmt.Println(c) // [true false false true]

	//	二维数组
	d := [...][3]string{
		{"aaa1", "bbb2", "ccc3"},
		{"aaa4", "bbb5", "ccc6"},
		{"aaa7", "bbb8", "ccc9"},
	}
	for _, value1 := range d {
		fmt.Println(value1)
		for i := 0; i < len(value1); i++ {
			fmt.Println(value1[i])
		}
	}
	//[aaa1 bbb2 ccc3]
	//aaa1
	//bbb2
	//ccc3
	//[aaa4 bbb5 ccc6]
	//aaa4
	//bbb5
	//ccc6
	//[aaa7 bbb8 ccc9]
	//aaa7
	//bbb8
	//ccc9

	//
	arr := [...]int{4, 5, 7, 1, 9}
	for i1, v1 := range arr {
		if v1 < 8 {
			for i2, v2 := range arr {
				if v1+v2 == 8 {
					fmt.Println(i1, i2)
				}
			}
		}
	}

}
