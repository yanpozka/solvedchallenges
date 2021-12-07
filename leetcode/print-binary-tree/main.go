package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	{
		r := &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
		}
		// fmt.Println(ht(r))
		m := printTree(r)
		for _, line := range m {
			fmt.Println(line)
		}
	}
	{
		r := &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 1},
		}
		m := printTree(r)
		for _, line := range m {
			fmt.Println(line)
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	h := ht(root)
	m := make([][]string, h)
	cols := int(math.Pow(2, float64(h+1))) - 1
	// fmt.Println("cols", cols)
	for ix := range m {
		m[ix] = make([]string, cols)
	}
	type coord struct {
		r, c int
		t    *TreeNode
	}
	q := []*coord{{r: 0, c: (cols - 1) / 2, t: root}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		t := curr.t
		m[curr.r][curr.c] = strconv.Itoa(t.Val)

		if t.Left != nil {
			c := curr.c - int(math.Pow(2, float64(h-curr.r-1)))
			q = append(q, &coord{r: curr.r + 1, c: c, t: t.Left})
		}
		if t.Right != nil {
			c := curr.c + int(math.Pow(2, float64(h-curr.r-1)))
			q = append(q, &coord{r: curr.r + 1, c: c, t: t.Right})
		}
	}

	return m
}

func ht(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(ht(root.Left), ht(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
