//
// https://leetcode.com/problems/longest-word-in-dictionary/description/
// https://play.golang.org/p/DRKQ2mgfnAV
package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []string{"w", "wo", "wor", "worl", "world"}
	fmt.Println("longest word:", longestWord(in))

	in = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println("longest word:", longestWord(in))

	in = []string{"ts", "e", "x", "pbhj", "opto", "xhigy", "erikz", "pbh", "opt", "erikzb", "eri", "erik", "xlye", "xhig", "optoj", "optoje", "xly", "pb", "xhi", "x", "o"}
	fmt.Println("longest word:", longestWord(in))

}

func longestWord(words []string) string {
	var m [31][]string
	dict := map[string]bool{}
	var maxLen int

	for _, w := range words {
		m[len(w)] = append(m[len(w)], w)
		dict[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	for ix := maxLen; ix > 0; ix-- {
		ws := m[ix]
		if len(ws) == 0 {
			continue
		}
		sort.Strings(ws)
		for _, w := range ws {
			var ok bool
			for k, end := 1, len(w); k < end; k++ {
				if dict[w[:k]] {
					ok = true
				} else {
					ok = false
					break
				}
			}
			if ok {
				return w
			}
			if ix == 1 {
				return w
			}
		}
	}
	return ""
}
