package bitwise

import "testing"

func TestBitwise(t *testing.T) {
	if rangeBitwiseAnd(5, 7) != 4 {
		t.Fatal("expected 4")
	}
	if rangeBitwiseAnd(0, 1) != 0 {
		t.Fatal("expected 0")
	}
	if rangeBitwiseAnd(5, 17) != 0 {
		t.Fatal("expected 4")
	}
}
