package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	var N, Q int

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = line[0 : len(line)-1]
	parts := strings.Split(line, " ")
	N, _ = strconv.Atoi(parts[0])
	Q, _ = strconv.Atoi(parts[1])

	tree = make([]int, N+1)
	rank = make([]int, N+1)
	population = make([]int, N+1)

	for ix := range tree {
		tree[ix] = ix
		population[ix] = 1
	}

	for ; Q > 0; Q-- {
		line, _ := reader.ReadString('\n')
		line = line[0 : len(line)-1]
		parts := strings.Split(line, " ")

		elem, _ := strconv.Atoi(parts[1])

		if parts[0] == "M" {
			b, _ := strconv.Atoi(parts[2])
			// fmt.Println("Merging", elem, b)
			merge(elem, b)
			// fmt.Println(tree, population)
		} else {
			fmt.Println(population[find(elem)])
		}
	}

}
