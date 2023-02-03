package main

import "fmt"

func main() {
	{
		t := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 9},
			},
		}
		fmt.Println(largestValues(t))
	}
	{
		t := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}
		fmt.Println(largestValues(t))
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	s := new(sol)
	s.largestValuesDFS(root, 0)

	return s.levels
}

type sol struct {
	levels []int
}

func (s *sol) largestValuesDFS(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(s.levels) == level {
		s.levels = append(s.levels, root.Val)
	} else {
		if root.Val > s.levels[level] {
			s.levels[level] = root.Val
		}
	}

	s.largestValuesDFS(root.Left, level+1)
	s.largestValuesDFS(root.Right, level+1)
}

func largestValuesBFS(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	type level struct {
		t *TreeNode
		l int
	}

	q := []level{{t: root, l: 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if len(result) == curr.l {
			result = append(result, curr.t.Val)
		} else if curr.t.Val > result[curr.l] {
			result[curr.l] = curr.t.Val
		}

		if curr.t.Left != nil {
			q = append(q, level{t: curr.t.Left, l: curr.l + 1})
		}

		if curr.t.Right != nil {
			q = append(q, level{t: curr.t.Right, l: curr.l + 1})
		}
	}

	return result
}
