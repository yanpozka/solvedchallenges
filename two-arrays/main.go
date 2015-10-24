package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		N, _ := strconv.Atoi(scanner.Text())

		A := make([]int, N)
		B := make([]int, N)

		scanner.Scan()
		K, _ := strconv.Atoi(scanner.Text())

		for ix := 0; ix < N && scanner.Scan(); ix++ {
			A[ix], _ = strconv.Atoi(scanner.Text())
		}
		for ix := 0; ix < N && scanner.Scan(); ix++ {
			B[ix], _ = strconv.Atoi(scanner.Text())
		}

		sort.Ints(A)
		sort.Sort(sort.Reverse(sort.IntSlice(B)))

		var ok bool = true
		for ix := 0; ix < N && ok; ix++ {
			if A[ix]+B[ix] < K {
				ok = false
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
