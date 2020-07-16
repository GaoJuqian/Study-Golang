package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	a := make(map[string]string, 9)
	a["name"] = "value1"
	a["age"] = "value2"

	fmt.Printf("%#v\n", a)

	v, ok := a["sex"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}
	// map是无序的
	for key, value := range a {
		fmt.Println(key, value)
	}
	delete(a, "name")
	fmt.Printf("%#v\n", a)

	//	map排序
	b := make(map[string]int, 100)
	// 添加map元素
	for i := 1; i <= 50; i++ {
		name := fmt.Sprintf("stu%02d", i) // 两位数格式化
		b[name] = rand.Intn(100)          // 随机数 100以内
	}
	// 无序map
	//for k, v := range b {
	//	fmt.Println(k, v)
	//}

	// 创建切片
	bSlice := make([]string, 0, 100)
	// 添加切片元素
	for k, _ := range b {
		bSlice = append(bSlice, k)
	}
	// 排序
	sort.Strings(bSlice)
	// 根据排序的切片 遍历map
	for _, v := range bSlice {
		fmt.Println(v, b[v])
	}

	// map 切片
	mapSlice := make(map[string][]int, 10)
	fmt.Println(mapSlice)
	mapSlice["asdasd"] = make([]int, 10)
	fmt.Println(mapSlice)
	//	切片map
	sliceMap := make([]map[string]int, 10)
	fmt.Println(sliceMap)
	sliceMap[0] = make(map[string]int, 10)
	sliceMap[0]["aaa"] = 1
	fmt.Println(sliceMap)

	//	小练习
	//	统计单词出现的次数
	var str = "how do you do how do you do how do you do how do you do"
	str2 := strings.Split(str, " ")
	fmt.Println(str2)
	strMap := make(map[string]int)
	for _, v := range str2 {
		strMap[v] += 1
	}
	fmt.Println(strMap)
	// 写出输出
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
