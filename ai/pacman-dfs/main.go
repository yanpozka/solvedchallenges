package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	WALL  = 37 // %
	EMPTY = 45 // -
)

var pac_row, pac_col, food_row, food_col, N, M int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanf("%d%d", &pac_row, &pac_col)
	fmt.Scanf("%d%d", &food_row, &food_col)
	fmt.Scanf("%d%d", &N, &M)

	Matrix := make([][]uint8, N)

	for i := 0; i < N && scanner.Scan(); i++ {
		var line string = scanner.Text()
		Matrix[i] = make([]uint8, M)

		for k := range line {
			Matrix[i][k] = line[k]
		}
	}
	fmt.Println(Matrix)
}
