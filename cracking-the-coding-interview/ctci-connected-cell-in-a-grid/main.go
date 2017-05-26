package main

import "fmt"

var n, m int

func main() {
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for r := range matrix {
		matrix[r] = make([]int, m)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			fmt.Scan(&matrix[r][c])
		}
	}
	var max int
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if matrix[r][c] == 1 {
				if g := dfs(r, c, matrix); g > max {
					max = g
				}
			}
		}
	}

	fmt.Println(max)
}

func dfs(r, c int, mtx [][]int) int {
	mtx[r][c] = 2

	count := 1
	if n_r, n_c := r-1, c-1; n_r >= 0 && n_c >= 0 && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r-1, c; n_r >= 0 && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r-1, c+1; n_r >= 0 && n_c < m && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r, c+1; n_c < m && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}

	if n_r, n_c := r+1, c+1; n_c < m && n_r < n && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r+1, c; n_r < n && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r+1, c-1; n_c >= 0 && n_r < n && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}
	if n_r, n_c := r, c-1; n_c >= 0 && mtx[n_r][n_c] == 1 {
		count += dfs(n_r, n_c, mtx)
	}

	return count
}
