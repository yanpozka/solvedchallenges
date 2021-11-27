package main

import "fmt"

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
	fmt.Println(numberOfSubstrings("aaacb"))
	fmt.Println(numberOfSubstrings("abc"))
}

func numberOfSubstrings(s string) int {
	var count [3]int
	var a, b, sol int

	for b < len(s) {
		count[s[b]-'a']++
		for count[0] > 0 && count[1] > 0 && count[2] > 0 && a <= b {
			count[s[a]-'a']--
			sol += len(s) - b
			a++
		}
		b++
	}

	return sol
}
