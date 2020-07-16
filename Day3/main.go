package main

import "fmt"

func main() {
	// 位运算符 & 判断奇偶数
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i&1 == 1) // 奇数
	//	fmt.Println(i&1 == 2) // 偶数
	//}

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", j, i, i*j)
		}
		fmt.Println()
	}

}
