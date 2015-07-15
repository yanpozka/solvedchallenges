package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const INF = int(MaxUint >> 1)

func main() {
	var N, M, q, x, y, w int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	N, _ = strconv.Atoi(parts[0])
	M, _ = strconv.Atoi(parts[1])
	N += 1 // !!

	Graph := make([][]int, N)

	for ix := 1; ix < N; ix++ {
		Graph[ix] = make([]int, N)

		for k := 1; k < N; k++ {
			Graph[ix][k] = INF
		}
	}

	for ; M > 0 && scanner.Scan(); M-- {
		parts := strings.Split(scanner.Text(), " ")
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
		w, _ = strconv.Atoi(parts[2])

		Graph[x][y] = w
	}

	for u := 1; u < N; u++ { // pivot u has to be the first
		for a := 1; a < N; a++ {
			if a == u {
				continue
			}
			for b := 1; b < N; b++ {
				if b == a || Graph[a][u] == INF || Graph[u][b] == INF {
					continue
				}
				var sum int = Graph[a][u] + Graph[u][b]
				if sum < Graph[a][b] {
					Graph[a][b] = sum
				}
			}
		}
	}
	scanner.Scan()
	q, _ = strconv.Atoi(scanner.Text())

	for ; q > 0 && scanner.Scan(); q-- {
		parts := strings.Split(scanner.Text(), " ")
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])

		if x == y {
			fmt.Println(0)
			continue
		}
		if Graph[x][y] != INF {
			fmt.Printf("%d\n", Graph[x][y])
		} else {
			fmt.Printf("%d\n", -1)
		}
	}
}
