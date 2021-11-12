package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println(findDuplicateSubtrees(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 5},
		},
	}))

	// for _, t := range trees {
	// fmt.Println(t)
	// }
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	return fmt.Sprintf("%d", t.Val)
}

type sol struct {
	hs map[string][]*TreeNode
}

func (s *sol) preOrderHashLessMemory(root *TreeNode) string {
	if root == nil {
		return "<nil>"
	}

	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%03d", root.Val))
	io.WriteString(h, s.preOrderHashLessMemory(root.Left))
	io.WriteString(h, s.preOrderHashLessMemory(root.Right))

	hash := fmt.Sprintf("%x", h.Sum(nil))
	s.hs[hash] = append(s.hs[hash], root)
	return hash
}

func (s *sol) preOrderFaster(root *TreeNode) string {
	if root == nil {
		return "<nil>"
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%03d", root.Val))
	b.WriteString(s.preOrderFaster(root.Left))
	b.WriteString(s.preOrderFaster(root.Right))

	hash := b.String()
	s.hs[hash] = append(s.hs[hash], root)
	return hash
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	s := &sol{hs: make(map[string][]*TreeNode)}

	s.preOrderHashLessMemory(root)

	var res []*TreeNode
	for _, trees := range s.hs {
		if len(trees) <= 1 {
			continue
		}
		res = append(res, trees[0])
	}
	return res
}
