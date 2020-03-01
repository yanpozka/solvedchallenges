package main

import (
	"fmt"
	"sort"
)

func main() {

	m := [][]int{
		{4, 3, 3, 1, 1},
		{1, 0, 0, -1, -1},
		{-2, -2, -2, -2, -3},
		{-2, -2, -2, -3, -3},
		{-3, -3, -3, -3, -3},
	}
	fmt.Println(countNegatives(m), " ==? 17")
	fmt.Println(countNegativesLazy(m), " ==? 17")

	m = [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(m), " ==? 8")

	m = [][]int{
		{3, 2},
		{1, 0},
	}
	fmt.Println(countNegatives(m), " ==? 0")

	m = [][]int{
		{1, -1},
		{-1, -1},
	}
	fmt.Println(countNegatives(m), " ==? 3")

	m = [][]int{{-1}}
	fmt.Println(countNegatives(m), " ==? 1")

	m = [][]int{{0}}
	fmt.Println(countNegatives(m), " ==? 0")
}

func countNegatives(grid [][]int) int {
	var count int

	for _, row := range grid {
		pos := sort.Search(len(row), func(i int) bool { return row[i] <= 0 })
		if pos == len(row) {
			continue
		}
		count += len(row) - pos

		for ix := pos; ix < len(row) && row[ix] == 0; ix++ {
			count--
		}
	}
	return count
}

func countNegativesLazy(grid [][]int) int {
	var count int
	for r := range grid {
		for c := len(grid[r]) - 1; c >= 0; c-- {
			if grid[r][c] >= 0 {
				break
			}
			count++
		}
	}
	return count
}
