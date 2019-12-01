// https://leetcode.com/problems/repeated-substring-pattern/
// https://play.golang.org/p/fZ3BE16hihB
package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeatedSubstringPattern("aaaaaaaaaa"))
	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))

	fmt.Println(repeatedSubstringPattern("babbabbabbabbab"))

	fmt.Println(repeatedSubstringPattern("aabaaba"))
}

func repeatedSubstringPattern(s string) bool {

	for p, half := 1, len(s)/2; p <= half; p++ {
		if len(s)%p != 0 {
			continue
		}
		prev := s[:p]
		//fmt.Println("prev", prev)
		var allEquals bool
		for ix := p; ix < len(s); ix += p {
			//fmt.Println(s[ix : ix+p])
			if prev == s[ix:ix+p] {
				allEquals = true
			} else {
				allEquals = false
				break
			}
		}
		if allEquals {
			return true
		}
	}
	return false
}
