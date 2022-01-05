package main

import "fmt"

func main() {
	fmt.Println(checkAlmostEquivalent("abcdeef", "abaaacc"))
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var fa [26]int
	var fb [26]int
	for _, ch := range word1 {
		fa[ch-'a']++
	}
	for _, ch := range word2 {
		fb[ch-'a']++
	}

	for ix := range fa {
		if d := diff(fa[ix], fb[ix]); d > 3 {
			return false
		}
	}
	return true
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
