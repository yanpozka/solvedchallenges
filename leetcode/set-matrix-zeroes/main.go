package main

import (
	"fmt"
	"math"
)

func main() {
	{
		m := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		setZeroes(m)
		fmt.Println(m)
	}
	{
		m := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		setZeroes(m)
		fmt.Println(m)
	}
	{
		m := [][]int{
			{-1}, {2}, {3},
		}
		setZeroes(m)
		fmt.Println(m)
	}
}

func setZeroes(matrix [][]int) {
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == 0 {
				setRow(matrix, r)
				setCol(matrix, c)
			}
		}
	}

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == math.MaxInt64 {
				matrix[r][c] = 0
			}
		}
	}
}

func setRow(matrix [][]int, row int) {
	for c := range matrix[row] {
		if matrix[row][c] != 0 {
			matrix[row][c] = math.MaxInt64
		}
	}
}

func setCol(matrix [][]int, col int) {
	for r := range matrix {
		if matrix[r][col] != 0 {
			matrix[r][col] = math.MaxInt64
		}
	}
}
