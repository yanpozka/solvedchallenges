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
	N, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, N)

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		nums[ix], _ = strconv.Atoi(scanner.Text())
	}

	p := nums[0]
	for _, n := range nums {
		if n < p {
			fmt.Print(n, " ")
		}
	}
	fmt.Print(p, " ")
	for _, n := range nums {
		if n > p {
			fmt.Print(n, " ")
		}
	}
}
