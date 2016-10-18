package main

import (
	"fmt"
	"math/rand"
)

const (
	empty     = 45
	hit       = 104
	miss      = 109
	destroyed = 100
)

func main() {
	var n int
	fmt.Scan(&n)
	var line string

	var Matrix [10][10]byte

	for r := 0; r < 10; r++ {
		fmt.Scan(&line)
		for c := 0; c < 10; c++ {
			Matrix[r][c] = line[c]
		}
	}

	next_hit_pair := make(chan coord)
	// exit := make(chan struct{})
	go nextHit(Matrix, next_hit_pair)

	for p := range next_hit_pair {
		if p.row+1 < 10 && Matrix[p.row+1][p.col] == empty {
			fmt.Println(p.row+1, p.col)
			return
		}
		if p.row-1 >= 0 && Matrix[p.row-1][p.col] == empty {
			fmt.Println(p.row-1, p.col)
			return
		}
		if p.col+1 < 10 && Matrix[p.row][p.col+1] == empty {
			fmt.Println(p.row, p.col+1)
			return
		}
		if p.col-1 >= 0 && Matrix[p.row][p.col-1] == empty {
			fmt.Println(p.row, p.col-1)
			return
		}
	}

	var empties []coord
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			switch Matrix[r][c] {
			case empty:
				empties = append(empties, coord{row: r, col: c})
			}
		}
	}
	r := empties[rand.Intn(len(empties))]
	fmt.Println(r.row, r.col)
}

type coord struct {
	row, col int
}

func nextHit(matrix [10][10]byte, next_hit_pair chan coord) { // , exit chan struct{}
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			switch matrix[r][c] {
			case hit:
				next_hit_pair <- coord{row: r, col: c}
			}
		}
	}
	close(next_hit_pair)
}
