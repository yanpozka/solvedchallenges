package main

import "fmt"

func main() {
	fmt.Printf("%q\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%q\n", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	pref := map[string]int{}
	for _, word := range strs {
		for size := 1; size <= len(word); size++ {
			pref[word[:size]]++
		}
	}
	// fmt.Println(pref)
	var max int
	var maxPref string
	for p, freq := range pref {
		if freq < len(strs) {
			continue
		}
		if freq > max {
			max = freq
			maxPref = p
		} else if freq == max {
			if len(p) > len(maxPref) {
				maxPref = p
			}
		}
	}
	return maxPref
}
