package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())

	var inverse_row, sum, inverse_sum int = N - 1, 0, 0
	for row := 0; row < N && scanner.Scan(); row++ {

		parts := strings.Split(scanner.Text(), " ")

		for col, numstr := range parts {
			num, _ := strconv.Atoi(numstr)

			if col == row {
				sum += num
			}
			if col == inverse_row {
				inverse_sum += num
			}
		}
		inverse_row-- // !!
	}
	fmt.Println(math.Abs(float64(sum - inverse_sum)))
}
