//
// https://leetcode.com/problems/surrounded-regions/description/
// https://play.golang.org/p/FtScj6FFWjT
package main

import (
	"fmt"
)

func main() {
	{
		b := [][]byte{
			[]byte("XXXX"),
			[]byte("XOOX"),
			[]byte("XXOX"),
			[]byte("XOXX"),
		}
		fmt.Println(b)
		solve(b)
		fmt.Println(b)
	}
	{
		b := [][]byte{
			[]byte("XOX"),
			[]byte("OXO"),
			[]byte("XOX"),
		}
		// Input:    [["X","O","X"],["O","X","O"],["X","O","X"]]
		// Expected: [["X","O","X"],["O","X","O"],["X","O","X"]]
		fmt.Println(b)
		solve(b)
		fmt.Println(b)
	}
	{
		b := [][]byte{
			[]byte("XOXOXO"),
			[]byte("OXOXOX"),
			[]byte("XOXOXO"),
			[]byte("OXOXOX"),
		}
		// Input: [["X","O","X","O","X","O"],["O","X","O","X","O","X"],["X","O","X","O","X","O"],["O","X","O","X","O","X"]]
		fmt.Println(b)
		solve(b)
		fmt.Println(b)
	}
}

func solve(board [][]byte) {
	N := len(board)
	if N <= 2 {
		return
	}
	M := len(board[0])
	initSet(N * M)

	check := func(r, c int) int {
		ix := r*M + c
		if c < M-1 && board[r][c+1] == 'O' {
			ixR := r*M + (c + 1)
			merge(ix, ixR)
		}
		if r < N-1 && board[r+1][c] == 'O' {
			ixD := ((r + 1) * M) + c
			merge(ix, ixD)
		}
		return ix
	}

	visited := make([]int, 0, 1026)

	for r := 0; r < N; r++ {
		for c := 0; c < M; c++ {
			if board[r][c] != 'O' {
				continue
			}
			visited = append(visited, check(r, c))
		}
	}

	hasBorder := map[int]bool{}
	for c := 0; c < M; c++ {
		if board[0][c] == 'O' {
			hasBorder[find(c)] = true
		}
		if board[N-1][c] == 'O' {
			hasBorder[find((M*(N-1))+c)] = true
		}
	}

	for r := 0; r < N; r++ {
		if board[r][0] == 'O' {
			hasBorder[find(r*M)] = true
		}
		if board[r][M-1] == 'O' {
			hasBorder[find(r*M+(M-1))] = true
		}
	}

	for _, ix := range visited {
		if hasBorder[find(ix)] {
			continue
		}
		r, c := ix/M, ix%M
		board[r][c] = 'X'
	}
}

var tree []int
var rank []int

func initSet(n int) {
	tree = make([]int, n)
	for ix := range tree {
		tree[ix] = ix
	}
	rank = make([]int, n)
}

func find(x int) int {
	if tree[x] == x {
		return x
	}
	return find(tree[x])
}

func merge(x, y int) {
	pX, pY := find(x), find(y)
	if pX == pY {
		return
	}
	if rank[pX] > rank[pY] {
		tree[pY] = pX
	} else {
		if rank[pX] == rank[pY] {
			rank[pY]++
		}
		tree[pX] = pY
	}
}
