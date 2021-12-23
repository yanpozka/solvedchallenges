package main

import "fmt"

func main() {
	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
}

func firstPalindrome(words []string) string {
	for _, w := range words {
		if isPal(w) {
			return w
		}
	}
	return ""
}

func isPal(word string) bool {
	l, r := 0, len(word)-1
	for l < r {
		if word[l] != word[r] {
			return false
		}
		l++
		r--
	}
	return true
}
