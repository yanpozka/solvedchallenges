package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(customSortString("cba", "abcd"))
	fmt.Println(customSortString("cbafg", "abcd"))
}

func customSortString(order string, s string) string {
	var freqs [26]int
	for ix := range s {
		freqs[s[ix]-'a']++
	}
	var checked [26]bool

	var b strings.Builder
	for ix := range order {
		c := freqs[order[ix]-'a']
		str := strings.Repeat(string(order[ix]), c)
		b.WriteString(str)
		checked[order[ix]-'a'] = true
	}

	for ix := range checked {
		if checked[ix] {
			continue
		}
		c := freqs[ix]
		str := strings.Repeat(string('a'+ix), c)
		b.WriteString(str)
	}

	return b.String()
}
