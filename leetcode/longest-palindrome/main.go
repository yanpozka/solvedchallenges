package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("abccccdda"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) int {
	freqs := map[byte]int{}
	for ix := range s {
		freqs[s[ix]]++
	}

	// fmt.Println(freqs)
	var c, maxOddLen, oddNum int
	for _, f := range freqs {
		if f%2 == 0 {
			c += f
		} else {
			if f > maxOddLen {
				maxOddLen = f
				oddNum = 1
			} else if f == maxOddLen {
				oddNum++
			}
		}
	}

	c += maxOddLen
	if oddNum > 1 && maxOddLen > 1 {
		c += (oddNum - 1) * (maxOddLen - 1)
	}
	for _, f := range freqs {
		if f%2 != 0 && f != maxOddLen {
			c += f - 1
		}
	}
	return c
}
