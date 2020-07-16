package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [...]int{11, 22, 33, 44, 55}
	fmt.Println(a)
	//	 基于数组创建切片
	b := a[:]
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	c := b[2 : len(b)-1]
	fmt.Println(c)
	//	make创建
	d := make([]string, 5, 10) // 类型 长度 *容量
	fmt.Printf("%#v\n", d)
	fmt.Println(len(d))
	fmt.Println(cap(d)) // 查看容量
	//	切片不能直接对比 == ，可以 nil（引用类型的零值）

	var a1 []bool // nil
	if a1 == nil {
		fmt.Println("a1:", a1)
	}
	a2 := []bool{}
	if a2 == nil {
		fmt.Println("a2", a2)
	}
	a3 := make([]bool, 0, 0)
	if a3 == nil {
		fmt.Println("a3", a3)
	}

	//
	aa := [...]int{1, 2, 3}
	aa[0] = 2
	fmt.Println(aa)
	Change(aa)
	fmt.Println(aa)
	//[2 2 3]
	//[100 2 3]
	//[2 2 3]

	//	扩容
	slice1 := []int{}
	if slice1 != nil {
		fmt.Println(slice1, cap(slice1), len(slice1))
	}
	slice1 = append(slice1, 1)
	fmt.Println(slice1, cap(slice1), len(slice1))
	// 切片删除
	slice2 := []string{"北", "上", "广", "深"}
	slice2 = append(slice2[1:3], slice2[3:]...)
	fmt.Println(slice2, cap(slice2), len(slice2))

	// 扩容问题
	test1 := make([]int, 5, 10)
	fmt.Println(test1, len(test1), cap(test1))
	for i := 0; i < 16; i++ { // 翻倍扩容
		test1 = append(test1, i)
	}
	fmt.Println(test1, len(test1), cap(test1))

	//	练习题
	var lianxi_a = [...]int{3, 7, 8, 9, 1} // sort排序
	lianxi_b := lianxi_a[:]
	sort.Ints(lianxi_b)
	fmt.Println(lianxi_b)
	//	lianxi1() > [ "0","1",...,"9" ]
	lianxi1()
}
func lianxi1() { // 写出输出
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a, len(a), cap(a)) // [     0 1 2 3 4 5 6 7 8 9] 15 20
}
func Change(arr [3]int) {
	arr[0] = 100
	fmt.Println(arr)

}
