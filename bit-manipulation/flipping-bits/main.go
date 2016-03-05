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

	for ; T > 0 && scanner.Scan(); T-- {
		n64, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		n := uint32(n64)

		fmt.Println(^n)
	}
}
