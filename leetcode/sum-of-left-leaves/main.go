package main

import "fmt"

func main() {
	fmt.Println(sumOfLeftLeaves(&TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	if root.Left != nil {
		if isLeaf(root.Left) {
			sum = root.Left.Val
		} else {
			sum = sumOfLeftLeaves(root.Left)
		}
	}
	sum += sumOfLeftLeaves(root.Right)
	return sum
}

func isLeaf(r *TreeNode) bool {
	if r.Left == nil && r.Right == nil {
		return true
	}
	return false
}
