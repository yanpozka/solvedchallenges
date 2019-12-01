package nqueens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}

	set := make(map[string]bool, 1)
	b := makeBoard(n)
	used := make([][]bool, n)
	for ix := range used {
		used[ix] = make([]bool, n)
	}

	return nqueens(n, 0, used, b, set)
}

func makeBoard(n int) [][]byte {
	board := make([][]byte, n)
	for r := 0; r < n; r++ {
		board[r] = []byte(strings.Repeat(".", n))
	}
	return board
}

func nqueens(n, level int, used [][]bool, b [][]byte, set map[string]bool) [][]string {
	if level > n {
		return nil
	}
	if countQs(b) == n {
		res := make([]string, 0, len(b))

		var key string
		for ix := range b {
			res = append(res, string(b[ix]))
			key += string(b[ix])
		}
		if set[key] {
			return nil
		}
		set[key] = true
		return [][]string{res}
	}

	var sols [][]string

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if level > 0 && !checkQ(r, c, n, b) {
				continue
			}
			if used[r][c] {
				continue
			}

			dupB := copyB(b)
			dupB[r][c] = 'Q'
			possSol := nqueens(n, level+1, used, dupB, set)
			if possSol == nil || len(possSol) == 0 {
				continue
			}
			if level == 0 {
				used[r][c] = true
			}

			for _, s := range possSol {
				sols = append(sols, s)
			}
		}
	}

	return sols
}

// func printB(b [][]byte) {
// 	for ix := range b {
// 		fmt.Println(string(b[ix]))
// 	}
// 	fmt.Println()
// }

func countQs(b [][]byte) int {
	var count int
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b); c++ {
			if b[r][c] == 'Q' {
				count++
			}
		}
	}
	return count
}

func copyB(b [][]byte) [][]byte {
	r := make([][]byte, 0, len(b))

	for ix := range b {
		row := make([]byte, len(b[ix]))
		copy(row, b[ix])
		r = append(r, row)
	}
	return r
}

func checkQ(row, col, n int, b [][]byte) bool {
	if b[row][col] != '.' {
		return false
	}

	// colum
	for r, c := row, 0; c < n; c++ {
		if c != col && b[r][c] == 'Q' {
			return false
		}
	}
	// row
	for r, c := 0, col; r < n; r++ {
		if r != row && b[r][c] == 'Q' {
			return false
		}
	}

	// diagonal (right) \
	for r, c := row+1, col+1; r < n && c < n; r++ {
		if b[r][c] == 'Q' {
			return false
		}
		c++
	}

	// diagonal (right) /
	for r, c := row-1, col+1; r >= 0 && c < n; r-- {
		if b[r][c] == 'Q' {
			return false
		}
		c++
	}

	// diagonal (left) \
	for r, c := row-1, col-1; r >= 0 && c >= 0; r-- {
		if b[r][c] == 'Q' {
			return false
		}
		c--
	}

	// diagonal (left) /
	for r, c := row+1, col-1; r < n && c >= 0; r++ {
		if b[r][c] == 'Q' {
			return false
		}
		c--
	}

	return true
}
