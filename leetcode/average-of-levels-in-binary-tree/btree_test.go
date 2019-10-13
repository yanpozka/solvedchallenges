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

	result := averageOfLevels(root)

	expected := []float64{3.0, 14.5, 11.0}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}

func TestLevelOrderBottom01(t *testing.T) {
	root := &TreeNode{
		Val: 5,
	}
	result := averageOfLevels(root)
	expected := []float64{5.0}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}

	result = averageOfLevels(nil)
	expected = nil
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("result: %v != expected: %v", result, expected)
	}
}
