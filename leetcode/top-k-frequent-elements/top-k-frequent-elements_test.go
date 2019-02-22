package topkfrequentelements

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	{
		in := []int{1}
		k := 1
		out := []int{1}
		res := topKFrequent(in, k)

		if !reflect.DeepEqual(out, res) {
			t.Fatalf("expected output: %+v NOT EQUAL to result: %+v", out, res)
		}
	}
	{
		in := []int{2, 1, 1, 1, 2, 2, 3}
		k := 2
		out := []int{1, 2}
		res := topKFrequent(in, k)

		if !reflect.DeepEqual(out, res) {
			t.Fatalf("expected output: %+v NOT EQUAL to result: %+v", out, res)
		}
		t.Log(res)
	}
}
