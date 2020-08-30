package main

import "fmt"

type node struct {
	id       int
	name     string
	children []*node
}

//  广度优先搜索
func findNode(n *node, id int) *node {

	nodeList := make([]*node, 0)   // 初始化
	nodeList = append(nodeList, n) // 将第一个节点传入
	for len(nodeList) > 0 {        // 如果 nodeList 不为空 一直循环
		selfNode := nodeList[0] // 将当前第一个进行判断
		nodeList = nodeList[1:] // 将当前移除，剩余的继续循环
		if selfNode.id == id {  // 当前id return
			return selfNode
		}
		if len(selfNode.children) > 0 { // 如果当前包含子节点
			for _, children := range selfNode.children { // 将子节点加入 nodeList
				nodeList = append(nodeList, children)
			}
		}
	}
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
		name: "dasdad",
	}
	d := &node{
		id:   4,
		name: "dasdad",
	}
	c := &node{
		id:       3,
		name:     "dasdad",
		children: []*node{d, e, f},
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
	// 数据

	test1 := findNode(a, 5)
	fmt.Printf("找到了:%v\n", test1)
	/*
		&{1 dasdad [0xc000074240]}
		&{2 dasdad [0xc000074210]}
		&{3 dasdad [0xc0000741e0 0xc0000741b0 0xc000074180]}
		&{4 dasdad []}
		&{5 dasdad []}
		找到了:&{5 dasdad []}
	*/

}
