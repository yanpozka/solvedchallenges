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
	var Z = N*2 + 2

	tree = make([]int, Z)
	rank = make([]int, Z)
	population = make([]int, Z)

	for ix := range tree {
		tree[ix] = ix
		population[ix] = 1
	}

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())

		merge(a, b) //
	}

	var min, max int = 15000000, -1

	for ix, n := range population {
		if n <= 1 || tree[ix] != ix {
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println(min, max)
}

//
var tree, rank, population []int

//
func find(elem int) int {
	if tree[elem] == elem {
		return elem
	}
	return find(tree[elem])
}

//
func merge(a, b int) {
	var aP, bP int = find(a), find(b)
	if aP == bP {
		return
	}
	if rank[aP] > rank[bP] {
		tree[bP] = aP
		population[aP] += population[bP]
		population[bP] = population[aP]
	} else {
		tree[aP] = bP
		if rank[aP] == rank[bP] {
			rank[bP]++
		}

		population[bP] += population[aP]
		population[aP] = population[bP]
	}
}
