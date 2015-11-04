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

	elems := make([]int, N)

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		m, _ := strconv.Atoi(scanner.Text())
		elems[ix] = m
	}

	ix := N - 1
	last := elems[ix]
	ix--

	for ; ix >= 0 && elems[ix] > last; ix-- {

		for k := 0; k < N-1; k++ {
			if k == ix {
				fmt.Print(elems[k], elems[k], " ")
			} else {
				fmt.Print(elems[k], " ")
			}
		}
		fmt.Println()
	}
	ix++
	printArr(elems[:ix])
	fmt.Print(last, " ")
	printArr(elems[ix : N-1])
	fmt.Println()
}

func printArr(arr []int) {
	for _, e := range arr {
		fmt.Print(e, " ")
	}
}
