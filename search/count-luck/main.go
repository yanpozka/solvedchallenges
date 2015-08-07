package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0; T-- {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		M, _ := strconv.Atoi(scanner.Text())

		for ix := 0; ix < N && scanner.Scan(); ix++ {
		}
		scanner.Scan()
		K, _ := strconv.Atoi(scanner.Text())

	}
}
