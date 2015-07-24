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
	V, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		c, _ := strconv.Atoi(scanner.Text())

		if c >= V {
			fmt.Println(ix)
			return
		}
	}
	fmt.Println(N)
}
