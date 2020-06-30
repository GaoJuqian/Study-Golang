package main

import (
	"fmt"
	"strings"
)

//	字符串操作

func main() {
	// 字符串长度
	str1 := "abcdef"
	fmt.Println(len(str1))
	//	字符串拼接
	fmt.Println("aaa" + "bbb")
	str2 := fmt.Sprintf("%s - %s", "一", "二")
	fmt.Println(str2)
	//	字符串切割
	str3 := "a,b,c,d"
	fmt.Println(strings.Split(str3, ","))
	fmt.Printf("%T\n", strings.Split(str3, ","))
	//	判断是否包含
	fmt.Println(strings.Contains(str3, "a")) // true
	//	判断前缀
	fmt.Println(strings.HasPrefix(str3, "a")) // true
	//	判断后缀
	fmt.Println(strings.HasSuffix(str3, "a")) // false
	//	判断字串 首次出现
	fmt.Println(strings.Index(str3, "b")) // 2
	//	判断子串 最后出现
	fmt.Println(strings.LastIndex(str3, "a")) // 0
	//	join 对字符串切片连接
	fmt.Println(strings.Join(strings.Split(str3, ","), "-")) // a-b-c-d
	//
	str4 := "大大大aa"
	for a, b := range str4 { // for range 用来遍历带有中文的字符串
		fmt.Print(a)          // 下标
		fmt.Printf("%c\n", b) // 字符
	}

}
