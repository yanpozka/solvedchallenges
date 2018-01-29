//
// https://leetcode.com/problems/number-of-islands/description/
// https://play.golang.org/p/FwByaiX5JlI
package main

import "fmt"

func main() {
	g := [][]byte{
		// {1, 1, 1, 1, 0},
		// {1, 1, 0, 1, 0},
		// {1, 1, 0, 0, 0},
		// {0, 0, 0, 0, 0},
		[]byte("11110"),
		[]byte("11010"),
		[]byte("11000"),
		[]byte("00000"),
	}
	fmt.Println(numIslands(g))

	g = [][]byte{
		// {1, 1, 0, 0, 0},
		// {1, 1, 0, 0, 0},
		// {0, 0, 1, 0, 0},
		// {0, 0, 0, 1, 1},
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011"),
	}
	fmt.Println(numIslands(g))
}

func numIslands(grid [][]byte) int {
	visited := make(map[pair]bool, len(grid))
	var count int

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '0' {
				continue
			}
			p := pair{R: r, C: c}
			if visited[p] {
				continue
			}
			count++
			DFS(p, visited, grid)
		}
	}
	return count
}

type pair struct {
	R, C int
}

func DFS(p pair, visited map[pair]bool, grid [][]byte) {
	visited[p] = true

	check := func(r, c int) {
		if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r]) && grid[r][c] == '1' {
			pc := pair{R: r, C: c}
			if !visited[pc] {
				DFS(pc, visited, grid)
			}
		}
	}

	check(p.R-1, p.C)
	check(p.R, p.C+1)
	check(p.R+1, p.C)
	check(p.R, p.C-1)
}
