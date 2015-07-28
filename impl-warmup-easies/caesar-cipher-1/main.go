package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	lwz := rune("z"[0])
	upZ := rune("Z"[0])
	var K int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	var line string = scanner.Text()

	scanner.Scan()
	K, _ = strconv.Atoi(scanner.Text())
	K %= 26

	for _, c := range line {
		if !unicode.IsLetter(c) {
			fmt.Printf("%s", string(c))
			continue
		}
		char := c + rune(K)
		if unicode.IsLower(c) {
			if char > lwz {
				char = rune("a"[0]) + (char - lwz - 1)
			}
		} else if unicode.IsUpper(c) {
			if char > upZ {
				char = rune("A"[0]) + (char - upZ - 1)
			}
		}
		fmt.Printf("%s", string(char))
	}
	fmt.Println()
}
