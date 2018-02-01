//
// https://leetcode.com/problems/two-sum-iv-input-is-a-bst/description/
//
package main

import (
	"fmt"
	"sort"
)

func main() {
	t := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val:  -1,
			Left: &TreeNode{Val: -3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
	}

	fmt.Println(findTarget(t, -4))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	inL := inOrder(root)
	if len(inL) <= 1 {
		return false
	}

	for ixN, n := range inL {
		diff := k - n
		ix := sort.SearchInts(inL, diff)
		if ix == len(inL) || ix == ixN {
			continue
		}
		if inL[ix] == diff {
			fmt.Println(n, diff)
			return true
		}
	}

	return false
}

func inOrder(r *TreeNode) []int {
	var res []int
	if r == nil {
		return res
	}
	res = inOrder(r.Left)
	res = append(res, r.Val)
	res = append(res, inOrder(r.Right)...)
	return res
}
