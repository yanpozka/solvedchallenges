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

	printArr(quickSort(nums))
}

func quickSort(ar []int) []int {
	if len(ar) <= 1 {
		return ar
	}
	p := ar[0]
	a := make([]int, 0)
	b := make([]int, 0)

	for _, n := range ar {
		if n < p {
			a = append(a, n)
		}
	}
	a = quickSort(a)
	if len(a) > 1 {
		printArr(a)
	}
	a = append(a, p)

	for _, n := range ar {
		if n > p {
			b = append(b, n)
		}
	}
	b = quickSort(b)
	if len(b) > 1 {
		printArr(b)
	}

	a = append(a, b...)
	return a
}

func printArr(ar []int) {
	for ix, n := range ar {
		fmt.Print(n)
		if ix < len(ar)-1 {
			fmt.Print(" ")
		}
	}
	if len(ar) > 0 {
		fmt.Println()
	}
}
