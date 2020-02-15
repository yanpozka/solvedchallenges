package maxancestor

import "testing"

func TestMaxAncestorDiff(t *testing.T) {
	{
		tree := &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   6,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 7},
				},
			},
			Right: &TreeNode{
				Val: 10,
				Right: &TreeNode{
					Val:  14,
					Left: &TreeNode{Val: 13},
				},
			},
		}
		if out := maxAncestorDiff(tree); out != 7 {
			t.Fatalf("Expected 7, but got: %d", out)
		}
	}
	{
		tree := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		}
		if out := maxAncestorDiff(tree); out != 3 {
			t.Fatalf("Expected 3, but got: %d", out)
		}
	}
	{
		tree := &TreeNode{
			Val: 8,
		}
		if out := maxAncestorDiff(tree); out != 0 {
			t.Fatalf("Expected 0, but got: %d", out)
		}
	}
}
