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

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		counts[num]++
	}

	for n, c := range counts {
		for ; c > 0; c-- {
			fmt.Print(n, " ")
		}
	}
}
