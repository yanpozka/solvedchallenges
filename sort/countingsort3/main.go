package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var counts [100]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		num, _ := strconv.Atoi(scanner.Text())
		scanner.Scan() // word

		counts[num]++
	}

	for ix := 1; ix < 100; ix++ {
		counts[ix] += counts[ix-1]
	}

	for _, c := range counts {
		fmt.Print(c, " ")
	}
}
