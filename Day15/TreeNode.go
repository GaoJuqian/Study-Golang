package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	a := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	test := find(a)
	fmt.Println(test)

}

// 遍历二叉树深度
func find(Tree *TreeNode) int {

	if Tree == nil {
		return 0
	}
	left := find(Tree.Left)

	right := find(Tree.Right)
	if right > left {
		return right + 1
	}
	return left + 1
}
