package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var R, C int

func DFS(M [][]bool, rw, cl int) int {
	M[rw][cl] = false
	var acum int

	if rw > 0 && M[rw-1][cl] { // up
		acum += DFS(M, rw-1, cl)
	}

	if rw > 0 && cl < C-1 && M[rw-1][cl+1] { // up-right
		acum += DFS(M, rw-1, cl+1)
	}

	if cl < C-1 && M[rw][cl+1] { // right
		acum += DFS(M, rw, cl+1)
	}

	if rw < R-1 && cl < C-1 && M[rw+1][cl+1] { // right-down
		acum += DFS(M, rw+1, cl+1)
	}

	if rw < R-1 && M[rw+1][cl] { // down
		acum += DFS(M, rw+1, cl)
	}

	if rw < R-1 && cl > 0 && M[rw+1][cl-1] { // left-down
		acum += DFS(M, rw+1, cl-1)
	}

	if cl > 0 && M[rw][cl-1] { // left
		acum += DFS(M, rw, cl-1)
	}

	if rw > 0 && cl > 0 && M[rw-1][cl-1] { // up-left
		acum += DFS(M, rw-1, cl-1)
	}

	return 1 + acum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	R, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	C, _ = strconv.Atoi(scanner.Text())

	matrix := make([][]bool, R)

	for ix := 0; ix < R; ix++ {
		matrix[ix] = make([]bool, C)
		for k := 0; k < C && scanner.Scan(); k++ {
			if scanner.Text() == "1" {
				matrix[ix][k] = true
			}
		}
	}

	var max int = -1

	for ix := 0; ix < R; ix++ {
		for k := 0; k < C; k++ {
			if matrix[ix][k] {
				if c := DFS(matrix, ix, k); c > max {
					max = c
				}
			}
		}
	}

	fmt.Println(max)
}
