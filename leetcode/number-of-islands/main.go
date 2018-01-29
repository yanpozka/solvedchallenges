//
// https://leetcode.com/problems/number-of-islands/description/
// https://play.golang.org/p/038oESZ5_VM
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
	N := len(grid)
	if N == 0 {
		return 0
	}
	M := len(grid[0])

	total := N*M + 1
	initTree(total)
	nodesCap := 1026
	if total < nodesCap {
		nodesCap = total
	}
	nodes := make([]int, 0, nodesCap)

	check := func(r, c, index int) {
		if r < N && c < M && grid[r][c] == '1' {
			merge(index, r*M+c)
		}
	}

	for ix, row := -1, 0; row < N; row++ {
		for col := 0; col < M; col++ {
			ix++
			if grid[row][col] == '0' {
				continue
			}
			nodes = append(nodes, ix)
			check(row, col+1, ix)
			check(row+1, col, ix)
		}
	}

	var count int
	for _, x := range nodes {
		if x == find(x) {
			count++
		}
	}
	return count
}

var tree []int
var rank []int

func initTree(n int) {
	tree = make([]int, n)
	for ix := range tree {
		tree[ix] = ix
	}
	rank = make([]int, n)
}

func find(x int) int {
	for x != tree[x] {
		x = tree[x]
	}
	return x
}

func merge(x, y int) {
	xP, yP := find(x), find(y)
	if xP == yP {
		return
	}
	if rank[xP] > rank[yP] {
		tree[yP] = xP
	} else {
		if rank[xP] == rank[yP] {
			rank[yP]++
		}
		tree[xP] = yP
	}
}
