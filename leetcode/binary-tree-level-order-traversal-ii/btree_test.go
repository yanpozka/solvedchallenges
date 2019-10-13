package btree

import (
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {

	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	result := levelOrderBottom(root)

	expected := [][]int{
		{15, 7},
		{9, 20},
		{3},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}

func TestLevelOrderBottom01(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	result := levelOrderBottom(root)
	expected := [][]int{
		{5},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}

	result = levelOrderBottom(nil)
	expected = nil
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}
