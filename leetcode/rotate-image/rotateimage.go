package rotateimage

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	for ix, lim := 0, n-1; ix < lim; ix++ {
		// row: 0
		lastRT := matrix[0][n-1]
		for k := n - 1; k > 0; k-- {
			matrix[0][k] = matrix[0][k-1]
		}
		matrix[0][0] = matrix[1][0]

		// col: n-1
		lastRB := matrix[n-1][n-1]
		for r := n - 1; r > 1; r-- {
			matrix[r][n-1] = matrix[r-1][n-1]
		}
		matrix[1][n-1] = lastRT

		// row: n-1
		lastLB := matrix[n-1][0]
		for k := 0; k < n-1; k++ {
			matrix[n-1][k] = matrix[n-1][k+1]
		}
		matrix[n-1][n-2] = lastRB

		// col: 0
		for r := 1; r < n-1; r++ {
			matrix[r][0] = matrix[r+1][0]
		}
		matrix[n-2][0] = lastLB
		// printM(matrix)
	}

	if n > 3 {
		rotate(innerM(matrix))
	}
}

func printM(matrix [][]int) {
	fmt.Println()
	for r := 0; r < len(matrix); r++ {
		fmt.Println(matrix[r])
	}
	fmt.Println()
}

func innerM(m [][]int) [][]int {
	res := make([][]int, len(m)-2)
	for ix, lim, k := 1, len(m)-1, 0; ix < lim; ix++ {
		res[k] = m[ix][1:lim]
		k++
	}
	return res
}
