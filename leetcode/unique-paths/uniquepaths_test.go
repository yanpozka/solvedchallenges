package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {

	if uniquePaths(3, 2) != 3 {
		t.FailNow()
	}

	if uniquePaths(7, 3) != 28 {
		t.FailNow()
	}
}
