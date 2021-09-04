package intersection

import "testing"

func TestIntervalIntersection(t *testing.T) {

	A := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}

	t.Log(intervalIntersection(A, B))
}
