package uniquepaths

var board [101][101]int

func uniquePaths(m int, n int) int {
	initBoard(m, n)
	return board[m-1][n-1]
}

func initBoard(m, n int) {
	for r := 0; r < m; r++ {
		board[r][0] = 1
	}
	for c := 0; c < n; c++ {
		board[0][c] = 1
	}
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			board[r][c] = 0
		}
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if r-1 >= 0 {
				board[r][c] += board[r-1][c]
			}
			if c-1 >= 0 {
				board[r][c] += board[r][c-1]
			}
		}
	}
}
