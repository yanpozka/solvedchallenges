package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func minPathSum(grid [][]int) int {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if c-1 >= 0 {
				if r-1 >= 0 {
					grid[r][c] += min(grid[r-1][c], grid[r][c-1])
				} else {
					grid[r][c] += grid[r][c-1]
				}
			} else if r-1 >= 0 {
				if c-1 >= 0 {
					grid[r][c] += min(grid[r-1][c], grid[r][c-1])
				} else {
					grid[r][c] += grid[r-1][c]
				}
			}
		}
	}
	// fmt.Println(grid)
	return grid[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
