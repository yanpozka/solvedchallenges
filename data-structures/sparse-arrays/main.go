package main

import "fmt"

func main() {
	var N, Q int
	var line string

	fmt.Scan(&N)
	words := make(map[string]int)

	for ; N > 0; N-- {
		fmt.Scan(&line)
		words[line]++
	}

	fmt.Scan(&Q)
	for ; Q > 0; Q-- {
		fmt.Scan(&line)
		fmt.Println(words[line])
	}
}
