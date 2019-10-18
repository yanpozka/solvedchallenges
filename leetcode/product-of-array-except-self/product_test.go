package product

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {

	expectedResult := []int{24, 12, 8, 6}
	res := productExceptSelf([]int{1, 2, 3, 4})

	if !reflect.DeepEqual(expectedResult, res) {
		t.Fatalf("expectedResult %v != result %v", expectedResult, res)
	}

	expectedResult = []int{24, 12, 8, 6}
	res = productExceptSelf([]int{1, 2, 3, 4})

	if !reflect.DeepEqual(expectedResult, res) {
		t.Fatalf("expectedResult %v != result %v", expectedResult, res)
	}
}
