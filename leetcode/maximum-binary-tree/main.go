package main

import "fmt"

func main() {
	root := constructMaximumBinaryTree([]int{1, 3, 2})
	fmt.Println(root.Val, root.Left.Val, root.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	var max, maxIx int
	for ix, n := range nums {
		if n > max {
			max = n
			maxIx = ix
		}
	}
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:maxIx]),
		Right: constructMaximumBinaryTree(nums[maxIx+1:]),
	}
}
