package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type cosa struct {
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	coins := make([]int64, M)
	nways := make([]int64, N+1)

	for ix := 0; ix < M && scanner.Scan(); ix++ {
		x, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		coins[ix] = x
	}

	NN := int64(N)
	nways[0] = 1
	for ix := 0; ix < M; ix++ {
		c := coins[ix]
		for k := c; k <= NN; k++ {
			nways[k] += nways[k-c]
		}
	}

	fmt.Println(nways[N])
}
