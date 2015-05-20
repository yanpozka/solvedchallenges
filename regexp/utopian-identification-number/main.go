package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var regEx = regexp.MustCompile(`^[a-z]{0,3}[0-9]{2,8}[A-Z]{3,}`)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	scanner := bufio.NewScanner(os.Stdin)

	for ; N > 0 && scanner.Scan(); N-- {
		if regEx.Match(scanner.Bytes()) {
			fmt.Println("VALID")
		} else {
			fmt.Println("INVALID")
		}
	}

}
