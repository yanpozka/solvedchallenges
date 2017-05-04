package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var tree []int

//
func initTree() {
	for ix := 1; ix < 64; ix++ {
		tree[ix] = ix
	}
}

//
func find(elem int) int {
	if tree[elem] == elem {
		return elem
	}
	return find(tree[elem])
}

//
func merge(a, b int) {
	if aP, bP := find(a), find(b); aP != bP {
		tree[bP] = aP
	}
}

//
func countComponents() int {
	var count int
	for ix, n := range tree {
		if n == ix {
			count++
		}
	}
	return count
}

//
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	nums := make([]uint64, N)

	tree = make([]int, 64)

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		nums[ix], _ = strconv.ParseUint(scanner.Text(), 10, 64)
	}

	var result int = 64

	for ix, S := 1, int(math.Pow(2, float64(N))); ix < S; ix++ { // 2^N for each sub-set:
		sub_set := uint32(ix)

		initTree() // reset tree

		for k := 0; k < N; k++ { // for each number of current sub-set

			if sub_set&1 == 1 {
				current_num, prev := nums[k], -1

				for b := 0; b < 64; b++ {
					if current_num&1 == 1 {
						if prev == -1 {
							prev = k
						} else {
							merge(b, prev) // merge !!
							prev = b
						}
					}
					current_num >>= 1
				}
			}

			sub_set >>= 1
		}

		result += countComponents()
	}

	fmt.Println(result)
}
