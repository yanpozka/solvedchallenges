package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	words := make(map[string]int)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for ; N > 0 && scanner.Scan(); N-- {
		parts := strings.Split(scanner.Text(), " ")

		for _, w := range parts {
			if _, contains := words[w]; !contains {
				words[w] = 0
			}
			words[w] += 1
		}
	}

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		var count int

		w := scanner.Text()
		if a, contains := words[w]; contains {
			count += a
		}

		uk := w[:len(w)-2] + "s" + w[len(w)-1:]

		if b, contains := words[uk]; contains {
			count += b
		}

		fmt.Println(count)
	}
}
