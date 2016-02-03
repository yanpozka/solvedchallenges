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
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	if N == M {
		fmt.Println(0)
		return
	}
	nums := make([]int, M)

	for ix := 0; ix < M && scanner.Scan(); ix++ {
		nums[ix], _ = strconv.Atoi(scanner.Text())
	}

	sort.Sort(sort.IntSlice(nums))

	max := nums[0]
	for ix := 1; ix < M; ix++ {
		if h := (nums[ix] - nums[ix-1]) / 2; h > max {
			max = h
		}
	}

	if h := (N - 1) - nums[M-1]; h > max {
		max = h
	}

	fmt.Println(max)
}
