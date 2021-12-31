package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}, 3))
}

func kWeakestRows(mat [][]int, k int) []int {
	type pair struct {
		val, ix int
	}
	rows := make([]*pair, len(mat))
	for r := 0; r < len(mat); r++ {
		var sum int
		for _, s := range mat[r] {
			sum += s
		}
		rows[r] = &pair{val: sum, ix: r}
	}

	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].val < rows[j].val {
			return true
		}
		if rows[i].val < rows[j].val {
			if i < j {
				return true
			}
		}
		return false
	})

	sol := make([]int, k)
	for ix, p := range rows[:k] {
		sol[ix] = p.ix
	}
	return sol
}
