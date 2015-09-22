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

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	var DD = "d"[0]

	for ; N > 0 && scanner.Scan(); N-- {
		original := scanner.Text()
		lw := strings.ToLower(original)

		if strings.HasPrefix(lw, "hi ") {
			if len(lw) >= 4 && lw[3] != DD {
				fmt.Println(original)
			}
		}
	}
}
