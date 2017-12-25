package main

import "testing"

func TestIsNotSymmetric(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	if isSymmetric(tree) {
		t.Error("Binary Tree is NOT symmetric")
	}
}

func TestIsSymmetric(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	if !isSymmetric(tree) {
		t.Error("Binary Tree has to be symmetric")
	}
}

func TestIsSymmetric2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
		},
	}

	if !isSymmetric(tree) {
		t.Error("Binary Tree has to be symmetric")
	}
}
