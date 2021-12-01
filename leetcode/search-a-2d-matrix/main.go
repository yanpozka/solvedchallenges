package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 12))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 30))
}

// fast solution: O(log m) + O(log n)
func searchMatrix(matrix [][]int, target int) bool {

	rowIx := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] >= target })
	if rowIx < len(matrix) && matrix[rowIx][0] == target {
		return true
	}
	if rowIx > 0 {
		rowIx--
	}
	row := matrix[rowIx]
	fIx := sort.SearchInts(row, target)
	if fIx != len(row) && row[fIx] == target {
		return true
	}

	return false
}

// simple solution: O( n * log m )
func searchMatrixSimple(matrix [][]int, target int) bool {
	for _, row := range matrix {
		fIx := sort.SearchInts(row, target)
		if fIx != len(row) && row[fIx] == target {
			return true
		}
	}
	return false
}
