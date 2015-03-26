package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var N int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	var nums []float64 = make([]float64, N)

	for ix := range parts {
		nums[ix], _ = strconv.ParseFloat(parts[ix], 32)
	}

	fmt.Println(avg(nums))

	fmt.Println(median(nums))
}

func avg(nums []float64) float64 {
	var total float64 = 0
	for i := 0; i < N; i++ {
		total += nums[i]
	}
	return total / float64(N)
}

func median(nums []float64) float64 {
	var floatSlice sort.Float64Slice = nums
	floatSlice.Sort()
	lenslice := len(floatSlice)

	if lenslice%2 != 0 {
		return floatSlice[lenslice/2.0]
	}

	return (floatSlice[lenslice/2] + floatSlice[lenslice/2-1]) / 2.0
}
