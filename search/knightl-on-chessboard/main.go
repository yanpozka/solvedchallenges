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
			// fmt.Print(search(n, r, c), " ")
			search(n, r, c)
		}
		// fmt.Println()
	}
}

func search(n, s, e int) int {
	b := make([][]bool, n)
	for ix := range b {
		b[ix] = make([]bool, n)
	}

	res := dfs(b, n, s, e, 0)
	for r := 1; r < n; r++ {
		for c := 1; c < n; c++ {
			fmt.Print(b[r][c], " ")
		}
		fmt.Println()
	}

	return res
}

func dfs(board [][]bool, n, r, c, count int) int {
	if r == n-1 && c == n-1 {
		return count + 1
	}
	fmt.Println(r, c, count)

	board[r][c] = true
	a, b := r, c

	// r +/- a
	// c +/- b
	if n_r, n_c := r+a, c+b; n_r < n && n_c < n && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r+a, c-b; n_r < n && n_c >= 0 && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r-a, c+b; n_r >= 0 && n_c < n && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r-a, c-b; n_r >= 0 && n_c >= 0 && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	// r +/- b
	// c +/- a
	if n_r, n_c := r+b, c+a; n_r < n && n_c < n && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r+b, c-a; n_r < n && n_c >= 0 && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r-b, c+a; n_r >= 0 && n_c < n && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}
	if n_r, n_c := r-b, c-a; n_r >= 0 && n_c >= 0 && board[n_r][n_c] == false {
		if result := dfs(duplicate(board), n, n_r, n_c, count+1); result != -1 {
			return result
		}
	}

	return -1
}

func duplicate(b [][]bool) [][]bool {
	return b

	// new_b := make([][]bool, len(b))
	// for r := range b {
	// 	new_b[r] = make([]bool, len(b))
	// 	for c := range b {
	// 		new_b[r][c] = b[r][c]
	// 	}
	// }
	// return new_b
}
