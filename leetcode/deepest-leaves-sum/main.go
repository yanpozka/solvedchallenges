package main

import "fmt"

func main() {
	fmt.Println(deepestLeavesSum(&TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	s := &sol{max: -1}
	s.heights(root, 0)
	return s.sum
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

type sol struct {
	max, sum int
}

func (s *sol) heights(root *TreeNode, prev int) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		if h := prev + 1; h > s.max {
			s.max = h
			s.sum = root.Val
		} else if h == s.max {
			s.sum += root.Val
		}
		return
	}

	s.heights(root.Left, prev+1)
	s.heights(root.Right, prev+1)
}
