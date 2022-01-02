package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("vim-go"))
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	bs := []byte(s)
	for l < r {
		if isVowel(bs[l]) && isVowel(bs[r]) {
			bs[l], bs[r] = bs[r], bs[l]
			l++
			r--
		} else {
			if !isVowel(bs[l]) {
				l++
			}
			if !isVowel(bs[r]) {
				r--
			}
		}
	}
	return string(bs)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
