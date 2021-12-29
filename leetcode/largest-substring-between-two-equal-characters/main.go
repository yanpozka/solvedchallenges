package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("bb"))
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
	fmt.Println(maxLengthBetweenEqualCharacters("cbzxy"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	type pair struct {
		f, l int
	}
	var chars [26]*pair
	for ix := range s {
		ch := s[ix] - 'a'
		if chars[ch] == nil {
			chars[ch] = &pair{f: ix}
		} else {
			chars[ch].l = ix
		}
	}

	var max int
	for _, ch := range chars {
		if ch == nil || ch.l == 0 {
			continue
		}
		if diff := ch.l - ch.f; diff > max {
			max = diff
		}
	}
	return max - 1
}
