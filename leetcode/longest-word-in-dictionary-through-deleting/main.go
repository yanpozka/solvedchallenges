package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}))
	fmt.Println(findLongestWord("abpcplea", []string{"x", "y", "z"}))
	fmt.Println(findLongestWord("aewfafwafjlwajflwajflwafj", []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}))
}

func findLongestWord(s string, dictionary []string) string {
	var freq [26][]int
	for ix, ch := range s {
		pos := ch - 'a'
		freq[pos] = append(freq[pos], ix)
	}

	var max string
	for _, word := range dictionary {
		if word == s {
			return word
		}
		if canBuild(freq, word) {
			// println("candidate:", word)
			if len(word) > len(max) {
				max = word
			} else if len(word) == len(max) {
				if word < max {
					max = word
				}
			}
		}
	}
	return max
}

func canBuild(freq [26][]int, word string) bool {
	// TODO: BFS here
	for ix, ch := range word {
		pos := ch - 'a'
		indices := freq[pos]
		if len(indices) == 0 {
			return false
		}
		foundIx := sort.SearchInts(indices, ix)
		if foundIx == len(indices) {
			return false
		}
		if indices[foundIx] < ix {
			return false
		}

		// add possible indices indices[foundIx:]

	}
	return true
}
