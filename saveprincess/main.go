package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, princess_row, princess_col int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < N && scanner.Scan(); i++ {
		var row string = scanner.Text()

		if i == 0 || i == N-1 { // first or last row
			var size int = len(row)
			if row[0] == 112 { // "p"
				princess_row, princess_col = i, 0
				break
			} else if row[size-1] == 112 { // "p"
				princess_row, princess_col = i, size-1
				break
			}
		}
	}
	var first_direct, second_direct string

	if princess_row == 0 {
		first_direct = "UP"
	} else {
		first_direct = "DOWN"
	}
	if princess_col == 0 {
		second_direct = "LEFT"
	} else {
		second_direct = "RIGHT"
	}
	var half int = N / 2
	for i := 0; i < half; i++ {
		fmt.Println(first_direct)
	}

	for i := 0; i < half; i++ {
		fmt.Println(second_direct)
	}
}
