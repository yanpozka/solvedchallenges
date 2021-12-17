package main

import "fmt"

func main() {
	t := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 8},
	}
	insertIntoBST(t, 2)
	fmt.Println(t.Val, t.Left.Val, t.Left.Left.Val, t.Left.Right)
	insertIntoBST(t, 4)
	fmt.Println(t.Val, t.Left.Val, t.Left.Left.Val, t.Left.Right)

	t1 := insertIntoBST(nil, 5)
	fmt.Println(t1.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	tmp := root
	for tmp != nil {
		if val < tmp.Val {
			if tmp.Left == nil {
				tmp.Left = &TreeNode{Val: val}
				break
			}
			tmp = tmp.Left
		} else {
			if tmp.Right == nil {
				tmp.Right = &TreeNode{Val: val}
				break
			}
			tmp = tmp.Right
		}
	}
	return root
}
