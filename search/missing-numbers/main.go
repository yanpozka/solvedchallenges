package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	countsA [103]int
	countsB [103]int
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var min int = 100000000

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	A := make([]int, N)
	for ix := 0; ix < N && scanner.Scan(); ix++ {
		num, _ := strconv.Atoi(scanner.Text())
		A[ix] = num
		if num < min {
			min = num
		}
	}

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())
	B := make([]int, M)
	for ix := 0; ix < M && scanner.Scan(); ix++ {
		num, _ := strconv.Atoi(scanner.Text())
		B[ix] = num
		if num < min {
			min = num
		}
	}

	for _, elem := range A {
		countsA[elem-min]++
	}
	for _, elem := range B {
		countsB[elem-min]++
	}

	H, L := countsA, countsB
	if M > N {
		H, L = countsB, countsA
	}

	for ix := 0; ix < 103; ix++ {
		if H[ix] != 0 {
			c := H[ix] - L[ix]
			n := ix + min
			if c >= 1 {
				fmt.Print(n, " ")
			}
		}
	}
}
