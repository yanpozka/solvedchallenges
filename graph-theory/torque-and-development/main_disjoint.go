package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func main() {
	var q, n, m, cl, cr int
	scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &q)

	for ; q > 0; q-- {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d%d%d%d", &n, &m, &cl, &cr)

		solution(n, m, cl, cr)
	}
}

func solution(n, m, c_lib, c_road int) {
	if c_road > c_lib == false {
		tree = make([]int, n+1)
		rank = make([]int, n+1)
		population = make([]int, n+1)

		for ix := range tree {
			tree[ix] = ix
			population[ix] = 1
		}
	}

	var s, d int
	for ; m > 0; m-- {
		scanner.Scan()
		bb := scanner.Bytes()

		if c_road > c_lib == false {
			s, d = 0, 0
			var ix int
			for _, v := range bb {
				ix++
				if v == ' ' {
					break
				}
				s = s*10 + int(v-'0')
			}

			for ; ix < len(bb); ix++ {
				d = d*10 + int(bb[ix]-'0')
			}

			merge(d, s)
		}
	}

	if c_road > c_lib {
		fmt.Println(n * c_lib)
		return
	}

	var result int64

	for ix := 1; ix <= n; ix++ {
		if find(ix) == ix {
			result += int64((population[ix]-1)*c_road + c_lib)
		}
	}

	fmt.Println(result)
}

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
