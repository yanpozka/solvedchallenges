package decodestring

import "testing"

func TestDecodeString(t *testing.T) {
	type testCase struct {
		in, expected string
	}

	for _, tc := range []*testCase{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"2[ab2[c]]3[cd]ef", "abccabcccdcdcdef"},
	} {
		if res := decodeString(tc.in); res != tc.expected {
			t.Fatalf("expected: '%s' but got: '%s'", tc.expected, res)
		}
	}
}
