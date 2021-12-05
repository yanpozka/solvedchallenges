package main

import "fmt"

func main() {
	{
		root := constructMaximumBinaryTree([]int{1, 3, 2})
		fmt.Println(root.Val, root.Left.Val, root.Right.Val)
	}
	{
		root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
		fmt.Println(root.Val, root.Left.Val, root.Right.Val)
	}
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

func constructMaximumBinaryTreeNonSlicing(nums []int) *TreeNode {
	return constructMaximumBT(nums, 0, len(nums)-1)
}

func constructMaximumBT(nums []int, start, end int) *TreeNode {
	if start == end {
		return &TreeNode{Val: nums[start]}
	}
	fmt.Println(start, end)
	if start > end {
		return nil
	}

	var max, maxIx int
	for ix := start; ix <= end; ix++ {
		if nums[ix] > max {
			max = nums[ix]
			maxIx = ix
		}
	}
	fmt.Println(max)
	return &TreeNode{
		Val:   max,
		Left:  constructMaximumBT(nums, start, maxIx-1),
		Right: constructMaximumBT(nums, maxIx+1, end),
	}
}
