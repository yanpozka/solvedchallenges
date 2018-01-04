package main

import "fmt"

func main() {
	t := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 7},
		},
	}

	po(t)
	deleteNode(t, 5)
	po(t)
}

func po(r *TreeNode) {
	_po(r)
	fmt.Println()
}

func _po(r *TreeNode) {
	if r == nil {
		return
	}
	fmt.Print(r.Val, " ")
	_po(r.Left)
	_po(r.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	}
	if root.Right == nil {
		return root.Left
	}
	cursor := root.Right
	for ; cursor.Left != nil; cursor = cursor.Left {
	}
	root.Val = cursor.Val
	root.Right = deleteNode(root.Right, cursor.Val)

	return root
}
