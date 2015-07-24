package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	var pos, neg, zeroes float32
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	for _, numstr := range parts {
		num, _ := strconv.Atoi(numstr)
		if num > 0 {
			pos += 1.0
		} else if num < 0 {
			neg += 1.0
		} else {
			zeroes += 1.0
		}
	}
	var Nfl float32 = float32(N)
	fmt.Printf("%.4f\n%.4f\n%.3f", pos/Nfl, neg/Nfl, zeroes/Nfl)
}
