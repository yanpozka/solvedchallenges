package main

import "fmt"

func main() {
	fmt.Println(minDiffInBST(&TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4},
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}
	inv := inOrder(root)
	if len(inv) == 1 {
		return 0
	}
	fmt.Println(inv)

	min := 999999999999
	for ix, size := 1, len(inv); ix < size; ix++ {
		if d := inv[ix] - inv[ix-1]; d < min {
			min = d
		}
	}

	return min
}

func inOrder(r *TreeNode) []int {
	if r == nil {
		return nil
	}

	res := inOrder(r.Left)
	res = append(res, r.Val)
	res = append(res, inOrder(r.Right)...)

	return res
}
