package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {

	if n := lengthOfLongestSubstring("dvdf"); n != 3 {
		t.Fatal(n, "dvdf")
	}

	// if n := lengthOfLongestSubstring(""); n != 0 {
	// 	t.Fatal(n, "")
	// }
	// if n := lengthOfLongestSubstring("abcabcbb"); n != 3 {
	// 	t.Fatal(n, "abcabcbb")
	// }
	// if n := lengthOfLongestSubstring("bbbbb"); n != 1 {
	// 	t.Fatal(n, "bbbbb")
	// }
	// if n := lengthOfLongestSubstring("pwwkew"); n != 3 {
	// 	t.Fatal(n, "bbbbb")
	// }
	// if n := lengthOfLongestSubstring("bbbba"); n != 2 {
	// 	t.Fatal(n, "bbbba")
	// }
	// if n := lengthOfLongestSubstring("b"); n != 1 {
	// 	t.Fatal(n, "b")
	// }
}
