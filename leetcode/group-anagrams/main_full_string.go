//
// https://leetcode.com/problems/group-anagrams/description/
// https://play.golang.org/p/UFKckQADN6q
package main

import (
	"bytes"
	"fmt"
)

func main() {
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(in))

	// In:       ["cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"]
	// Expected: [["doc"],["bar"],["buy"],["ill"],["may"],["tin"],["cab"],["pew"],["max"],["duh"]]

	in = []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(in))

	in = []string{"tao", "pit", "cam", "aid", "pro", "dog"}
	fmt.Println(groupAnagrams(in))

	fmt.Printf("%#v\n", groupAnagrams([]string{"", "cab", "", "", "abc"}))
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[int][]string, len(strs))

	for _, s := range strs {
		m[len(s)] = append(m[len(s)], s) // grouping by length
	}

	res := make([][]string, 0, len(m))
	for _, strs := range m {
		hashes := make(map[string][]string, len(strs))

		for _, s := range strs {
			hash := h(s)
			hashes[hash] = append(hashes[hash], s)
		}

		for _, strs := range hashes {
			res = append(res, strs)
		}
	}
	return res
}

func h(s string) string {
	var b bytes.Buffer
	if s == "" {
		return b.String()
	}

	var alpha [26]int
	for ix := 0; ix < len(s); ix++ {
		alpha[s[ix]-'a']++
	}

	for n, rep := range alpha {
		if rep == 0 {
			continue
		}
		b.Write(bytes.Repeat([]byte{byte('a' + n)}, rep))
	}
	return b.String()
}
