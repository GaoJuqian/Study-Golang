package main

import "fmt"

type node struct {
	id       int
	name     string
	children []*node
}

//  深度优先搜索
func findNode(n *node, id int) *node {
	if n.id == id {
		return n // 查找到，给当前递归返回
	}
	if len(n.children) > 0 {
		for _, v := range n.children {
			test := findNode(v, id) // 将上一层的返回取出
			fmt.Printf("当前所处:%v\n", n.id)
			if test != nil { // 取出的不为nil 将结果给下一层递归，直至出栈完毕
				return test
			}
		}
	}
	// 当没有子节点会 return nil，且没有查找到
	return nil
}

func main() {
	// 数据
	f := &node{
		id:   6,
		name: "dasdad",
	}
	e := &node{
		id:   5,
		name: "dasdad1",
	}
	d := &node{
		id:   4,
		name: "dasdad2",
	}
	c := &node{
		id:       3,
		name:     "dasdad",
		children: []*node{f, e, d}, // 展示深度的变化，颠倒了顺序

	}

	b := &node{
		id:       2,
		name:     "dasdad",
		children: []*node{c},
	}

	a := &node{
		id:       1,
		name:     "dasdad",
		children: []*node{b},
	}

	test1 := findNode(a, 5)
	fmt.Printf("找到了:%v\n", test1)

	//func a(i int) {
	//	fmt.Printf("前:%v\n", i)
	//	if i == 6 {
	//		//i
	//	} else {
	//		i++
	//		a(i)
	//	}
	//	fmt.Printf("后:%v\n", i)
	//}
	//	test := a(1)

	// 递归test
	//test := a(1)
	//fmt.Println(test)
}

func a(i int) int {
	fmt.Printf("前:%v\n", i)
	if i == 6 {
		return i
	} else {
		i++
		a := a(i)
		return a
	}
}
