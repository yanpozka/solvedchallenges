package btree

import (
	"reflect"
	"testing"
)

func TestIsCompleteTreeTrue(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}

	result := isCompleteTree(root)
	expected := true
	if result != expected {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}

func TestIsCompleteTreeFalse(t *testing.T) {

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}

	result := isCompleteTree(root)
	expected := false
	if result != expected {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}
func TestIsCompleteTree01(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	result := isCompleteTree(root)
	expected := true
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}

	result = isCompleteTree(nil)
	expected = true
	if result != expected {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}
