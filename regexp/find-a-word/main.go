package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regEx = regexp.MustCompile(`[[:word:]]+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	var buffer bytes.Buffer
	for ; N > 0; N-- {
		scanner.Scan()
		buffer.WriteString(scanner.Text())
		buffer.WriteString("\n")
	}

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	var text string = buffer.String()

	words_map := make(map[string]int)
	all_words := regEx.FindAllString(text, -1)

	for ix := range all_words {
		var current string = strings.ToLower(all_words[ix])
		if _, exists := words_map[current]; !exists {
			words_map[current] = 0
		}
		words_map[current]++
	}

	for ; T > 0; T-- {
		scanner.Scan()
		word := strings.ToLower(scanner.Text())
		var result int = 0
		if cr, exists := words_map[word]; exists {
			result = cr
		}
		fmt.Println(result)
	}
}
