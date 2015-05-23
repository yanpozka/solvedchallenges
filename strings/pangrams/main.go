package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	letters_count := make(map[byte]bool)

	scanner.Scan()
	var line string = strings.ToLower(scanner.Text())

	for ix := range line {
		if line[ix] >= 97 && line[ix] <= 122 { // a-z
			letters_count[line[ix]] = true
		}
	}

	if len(letters_count) != 26 {
		fmt.Println("not pangram")
	} else {
		fmt.Println("pangram")
	}
}
