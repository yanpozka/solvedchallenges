package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	solution(n)
}

func solution(n int) {

	for r := 1; r < n; r++ {
		for c := 1; c < n; c++ {
			fmt.Println(search(n, r, c))
		}
	}
}

func search(n, s, e int) int {
	b := make([][]bool, n)
	for ix := range b {
		b[ix] = make([]bool, n)
	}

	return dfs(b, n, s, e)
}

func dfs(b [][]bool, n, r, c, count int) int {
	if r == n-1 && c == n-1 {
		return count
	}

	b[r][c] = true

	// - -
	//    |
	if n_r, n_c := r+1, c+2; n_r < n && n_c < n && b[n_r][n_c] {
		if result := dfs(b, n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	// |
	// | _ >
	if n_r, n_c := r+2, c+1; n_r < n && n_c < n && b[n_r][n_c] {
		if result := dfs(b, n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	//     |
	// < _ |
	if n_r, n_c := r+1, c-2; n_r < n && n_c >= 0 && b[n_r][n_c] {
		if result := dfs(b, n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	//  _ _
	// |
	if n_r, n_c := r+1, c-2; n_r < n && n_c >= 0 && b[n_r][n_c] {
		if result := dfs(b, n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	// /\
	//  |
	//   - -
	if n_r, n_c := r-1, c-2; n_r >= 0 && n_c >= 0 && b[n_r][n_c] {
		if result := dfs(b, n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	return -1
}
