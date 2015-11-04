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

	for ix := 1; ix < N; ix++ {
		k := ix

		for ; k > 0 && elems[k-1] > elems[k]; k-- {
			elems[k-1], elems[k] = elems[k], elems[k-1]
		}
		printArr(elems, N)
	}
}

func printArr(arr []int, n int) {
	for ix, e := range arr {
		fmt.Print(e)
		if ix < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
