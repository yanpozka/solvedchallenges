package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func findWords(words []string) []string {
	fR := "qwertyuiop"
	sR := "asdfghjkl"
	tR := "zxcvbnm"
	var sol []string

	for _, w := range words {
		lw := strings.ToLower(w)
		if allInRow(lw, fR) {
			sol = append(sol, w)
		} else if allInRow(lw, sR) {
			sol = append(sol, w)
		} else if allInRow(lw, tR) {
			sol = append(sol, w)
		}
	}
	return sol
}

func allInRow(w, row string) bool {
	for ix := range w {
		if strings.IndexByte(row, w[ix]) == -1 {
			return false
		}
	}
	return true
}
