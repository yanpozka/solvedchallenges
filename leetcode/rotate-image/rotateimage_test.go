package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotateImageOneMatrix1n(t *testing.T) {
	input := [][]int{{4}}
	expected := [][]int{{4}}
	rotate(input)

	if reflect.DeepEqual(input, expected) == false {
		t.FailNow()
	}
}

func TestRotateImageOneMatrix2n(t *testing.T) {
	input := [][]int{
		{1, 2},
		{4, 3},
	}
	expected := [][]int{
		{4, 1},
		{3, 2},
	}
	rotate(input)

	if reflect.DeepEqual(input, expected) == false {
		t.FailNow()
	}
}

func TestRotateImageMatrix3n(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	rotate(input)

	if reflect.DeepEqual(input, expected) == false {
		t.FailNow()
	}
}

func TestRotateImageMatrix4n(t *testing.T) {
	input := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}
	rotate(input)

	if reflect.DeepEqual(input, expected) == false {
		t.FailNow()
	}
}
