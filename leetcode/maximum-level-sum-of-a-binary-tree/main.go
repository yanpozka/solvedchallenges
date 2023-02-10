package main

import "fmt"

func main() {
	fmt.Println(maxLevelSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}))
}

type sol struct {
	levels []int
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	s := new(sol)
	s.maxLevelSumDFS(root, 0)

	max, maxIx := s.levels[0], 0

	for ix := 1; ix < len(s.levels); ix++ {
		if s.levels[ix] > max {
			max, maxIx = s.levels[ix], ix
		}
	}

	return maxIx + 1
}

func (s *sol) maxLevelSumDFS(root *TreeNode, l int) {
	if root == nil {
		return
	}

	if len(s.levels) == l {
		s.levels = append(s.levels, root.Val)
	} else {
		s.levels[l] += root.Val
	}

	if root.Left != nil {
		s.maxLevelSumDFS(root.Left, l+1)
	}

	if root.Right != nil {
		s.maxLevelSumDFS(root.Right, l+1)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSumBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type lev struct {
		t *TreeNode
		l int
	}
	q := []*lev{{t: root, l: 0}}
	levels := []int{0}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if len(levels) == curr.l {
			levels = append(levels, curr.t.Val)
		} else {
			levels[curr.l] += curr.t.Val
		}

		if curr.t.Left != nil {
			q = append(q, &lev{t: curr.t.Left, l: curr.l + 1})
		}

		if curr.t.Right != nil {
			q = append(q, &lev{t: curr.t.Right, l: curr.l + 1})
		}
	}

	max, maxIx := levels[0], 0

	for ix := 1; ix < len(levels); ix++ {
		if levels[ix] > max {
			max, maxIx = levels[ix], ix
		}
	}

	return maxIx + 1
}
