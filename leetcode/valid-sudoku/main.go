package main

import "fmt"

func main() {
	b := [][]byte{
		{46, 56, 55, 54, 53, 52, 51, 50, 49},
		{50, 46, 46, 46, 46, 46, 46, 46, 46},
		{51, 46, 46, 46, 46, 46, 46, 46, 46},
		{52, 46, 46, 46, 46, 46, 46, 46, 46},
		{53, 46, 46, 46, 46, 46, 46, 46, 46},
		{54, 46, 46, 46, 46, 46, 46, 46, 46},
		{55, 46, 46, 46, 46, 46, 46, 46, 46},
		{56, 46, 46, 46, 46, 46, 46, 46, 46},
		{57, 46, 46, 46, 46, 46, 46, 46, 46},
	}
	fmt.Println(isValidSudoku(b))
}

func isValidSudoku(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		var checked [10]bool

		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			ix := board[r][c] - '0'
			if checked[ix] {
				return false
			}
			checked[ix] = true
		}
	}
	for c := 0; c < 9; c++ {
		var checked [10]bool

		for r := 0; r < 9; r++ {
			if board[r][c] == '.' {
				continue
			}
			ix := board[r][c] - '0'
			if checked[ix] {
				return false
			}
			checked[ix] = true
		}
	}

	checkQuad := func(sr, sc int) bool {
		var checked [10]bool
		for r := sr; r < sr+3; r++ {
			for c := sc; c < sc+3; c++ {
				if board[r][c] == '.' {
					continue
				}
				ix := board[r][c] - '0'
				if checked[ix] {
					return false
				}
				checked[ix] = true
			}
		}
		return true
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			if !checkQuad(r, c) {
				return false
			}
		}
	}

	return true
}
