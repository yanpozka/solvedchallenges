package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(countWords([]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}))
}

func countWords(words1 []string, words2 []string) int {
	f1 := map[string]int{}
	f2 := map[string]int{}

	for _, w := range words1 {
		f1[w]++
	}
	for _, w := range words2 {
		f2[w]++
	}

	var count int
	for w, c1 := range f1 {
		if c1 > 1 {
			continue
		}
		if c2 := f2[w]; c2 == 1 {
			count++
		}
	}
	return count
}

// in-place solution
func countWordsLessMemory(words1 []string, words2 []string) int {
	large, small := largeArr(words1, words2)
	sort.Strings(large)
	sort.Strings(small)

	var count int
	for ix := 0; ix < len(small); ix++ {
		if (ix-1 >= 0 && small[ix-1] == small[ix]) || (ix+1 < len(small) && small[ix+1] == small[ix]) {
			continue
		}
		w := small[ix]
		fIx := sort.SearchStrings(large, w)
		if fIx == len(large) {
			continue
		}
		if large[fIx] != w {
			continue
		}
		if fIx-1 >= 0 && large[fIx-1] == w {
			continue
		}
		if fIx+1 < len(large) && large[fIx+1] == w {
			continue
		}
		count++
	}
	return count
}

func largeArr(a, b []string) ([]string, []string) {
	if len(a) > len(b) {
		return a, b
	}
	return b, a
}
