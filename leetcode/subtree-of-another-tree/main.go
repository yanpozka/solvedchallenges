package main

import "fmt"

func main() {
	{
		r := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 5},
		}
		sr := &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}
		fmt.Println(isSubtree(r, sr))
	}
	{
		r := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val: 2, Left: &TreeNode{Val: 0},
				},
			},
			Right: &TreeNode{Val: 5},
		}
		sr := &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}
		fmt.Println(isSubtree(r, sr))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return root == nil
	}
	q := []*TreeNode{root}
	for len(q) != 0 {
		curr := q[0]
		q = q[1:] // pop

		if curr.Val == subRoot.Val {
			if equal(curr, subRoot) {
				return true
			}
		}
		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}

	return false
}

func equal(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
