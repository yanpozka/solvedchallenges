// https://leetcode.com/problems/find-common-characters/
// https://play.golang.org/p/kd-J_6Nk9iC
package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))

	fmt.Println(commonChars([]string{"bbddabab", "cbcddbdd", "bbcadcab", "dabcacad", "cddcacbc", "ccbdbcba", "cbddaccc", "accdcdbb"}))
}

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}

	var firstConts [26]int
	for ix := 0; ix < len(A[0]); ix++ {
		firstConts[A[0][ix]-'a']++
	}

	var res []string
	for ix := range firstConts {
		cont := firstConts[ix]
		if cont <= 0 {
			continue
		}
		first := byte('a' + ix)
		//fmt.Printf("%s: %d\n", string(first), cont)
		foundAll := true

		for k := 1; k < len(A); k++ {
			var found bool
			var minC int
			for c := 0; c < len(A[k]); c++ {
				if A[k][c] == first {
					minC++
					found = true
				}
			}
			if !found {
				foundAll = false
				break
			} else {
				if minC < cont {
					cont = minC
				}
			}
		}
		if foundAll {
			//fmt.Println(cont)
			for t := 0; t < cont; t++ {
				res = append(res, string(first))
			}
		}
	}
	return res
}
