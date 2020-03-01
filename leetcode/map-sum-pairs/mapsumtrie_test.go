package mapsumtrie

import "testing"

func TestMapSum(t *testing.T) {
	{
		mapSum := Constructor()

		mapSum.Insert("aa", 3)
		expected := 3
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("aa", 2)
		expected = 2
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("aaa", 2)
		expected = 4
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("aa", 1)
		expected = 3
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}
	}
	{
		mapSum := Constructor()

		mapSum.Insert("a", 3)
		expected := 0
		if s := mapSum.Sum("ap"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("b", 2)
		expected = 3
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}
	}
	{
		mapSum := Constructor()

		mapSum.Insert("apple", 3)
		// Input: sum("ap"), Output: 3
		expected := 3
		if s := mapSum.Sum("ap"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("app", 2)
		// Input: sum("ap"), Output: 5
		expected = 5
		if s := mapSum.Sum("ap"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("bab", 3)
		expected = 3
		if s := mapSum.Sum("ba"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("cat", 4)
		expected = 4
		if s := mapSum.Sum("cat"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}
		expected = 4
		if s := mapSum.Sum("c"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}
	}
	{
		mapSum := Constructor()

		mapSum.Insert("a", 1)
		expected := 1
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("ab", 1)
		expected = 1
		if s := mapSum.Sum("ab"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		mapSum.Insert("abc", 1)
		expected = 1
		if s := mapSum.Sum("abc"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		expected = 3
		if s := mapSum.Sum("a"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}

		expected = 2
		if s := mapSum.Sum("ab"); s != expected {
			t.Fatalf("expected sum: %d, but got: %d", expected, s)
		}
	}
}
