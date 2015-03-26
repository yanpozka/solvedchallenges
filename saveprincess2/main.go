package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, princess_row, princess_col int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	bot_row, _ := strconv.Atoi(parts[0])
	bot_col, _ := strconv.Atoi(parts[1])

	for i := 0; i < N && scanner.Scan(); i++ {
		var row string = scanner.Text()
		for k := 0; k < N; k++ {
			if row[k] == 112 { // "p"
				princess_row, princess_col = i, k
				break // we have the position of the princess so we don't need to continue read the input
			}
		}
	}
	var first_direct string

	if princess_col > bot_col {
		first_direct = "RIGHT"
	} else {
		if princess_col < bot_col {
			first_direct = "LEFT"
		} else { // same col
			if princess_row > bot_row {
				first_direct = "DOWN"
			} else {
				first_direct = "UP"
			}
		}
	}

	fmt.Println(first_direct)
}
