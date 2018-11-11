package uniquepathsii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	initBoard(obstacleGrid, m, n)

	return obstacleGrid[m-1][n-1]
}

func initBoard(board [][]int, m, n int) {
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == 1 {
				board[r][c] = -1
			}
		}
	}
	for r := 0; r < m && board[r][0] != -1; r++ {
		board[r][0] = 1
	}
	for c := 0; c < n && board[0][c] != -1; c++ {
		board[0][c] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			if board[r][c] == -1 {
				continue
			}

			if r-1 >= 0 && board[r-1][c] != -1 {
				board[r][c] += board[r-1][c]
			}
			if c-1 >= 0 && board[r][c-1] != -1 {
				board[r][c] += board[r][c-1]
			}
		}
	}
}
