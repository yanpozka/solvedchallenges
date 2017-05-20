package main

import (
	"container/list"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for r := 1; r < n; r++ {
		for c := 1; c < n; c++ {
			fmt.Print(bfs(n, r, c))
			if c < n-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func bfs(n, src_r, src_c int) int {
	board := make([][]bool, n)
	for ix := range board {
		board[ix] = make([]bool, n)
	}

	board[src_r][src_c] = true
	a, b := src_r, src_c

	q := list.New()
	q.PushBack(&pair{R: src_r, C: src_c, Level: 1})

	for q.Len() > 0 {
		p := q.Remove(q.Front()).(*pair)
		if p.R == n-1 && p.C == n-1 {
			return p.Level
		}
		r, c := p.R, p.C

		// r +/- a
		// c +/- b
		if n_r, n_c := r+a, c+b; n_r < n && n_c < n && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r+a, c-b; n_r < n && n_c >= 0 && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r-a, c+b; n_r >= 0 && n_c < n && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r-a, c-b; n_r >= 0 && n_c >= 0 && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}

		// r +/- b
		// c +/- a
		if n_r, n_c := r+b, c+a; n_r < n && n_c < n && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r+b, c-a; n_r < n && n_c >= 0 && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r-b, c+a; n_r >= 0 && n_c < n && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
		if n_r, n_c := r-b, c-a; n_r >= 0 && n_c >= 0 && board[n_r][n_c] == false {
			board[n_r][n_c] = true
			q.PushBack(&pair{R: n_r, C: n_c, Level: p.Level + 1})
		}
	}

	return -1
}

type pair struct {
	R, C, Level int
}
