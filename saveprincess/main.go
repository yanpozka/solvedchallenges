package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, pr_row, pr_col int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("p"[0])

	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())

	for i := 0; i < N && scanner.Scan(); i++ {
		var row string = scanner.Text()

		if i == 0 || i == N-1 { // first row
			var size int = len(row)
			if row[0] == 112 { // "p"
				pr_row, pr_col = i, 0
				break
			} else if row[size-1] == 112 {
				pr_row, pr_col = i, size-1
				break
			}
		}
	}
	fmt.Println(pr_row, pr_col)
}
