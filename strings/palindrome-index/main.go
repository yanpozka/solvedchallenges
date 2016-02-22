package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var line string
		fmt.Scan(&line)
		S := len(line)
		var result int = -1

		for a, b, H := 0, S-1, S/2; a < H; a++ {
			if line[a] != line[b] {
				if isPalindrome(line[:a] + line[a+1:]) {
					result = a
				} else {
					result = b
				}
				break
			}
			b--
		}
		fmt.Println(result)
	}
}

func isPalindrome(s string) bool {
	S := len(s)
	for a, b, H := 0, S-1, S/2; a < H; a++ {
		if s[a] != s[b] {
			return false
		}
		b--
	}
	return true
}