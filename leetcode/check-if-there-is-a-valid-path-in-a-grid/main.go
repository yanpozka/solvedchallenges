package main

import "fmt"

func main() {
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
	fmt.Println(hasValidPath([][]int{{1, 2, 1}, {1, 2, 1}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 2}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 1, 1, 1, 1, 3}}))
	fmt.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}, {2}, {6}}))
}

const (
	horizontal = iota + 1
	vertical
	leftDown
	rightDown
	leftUp
	rightUp
)

// BFS
func hasValidPath(grid [][]int) bool {
	if len(grid) == 0 {
		return false
	}
	n, m := len(grid), len(grid[0])
	type coord struct {
		r, c int
	}
	visited := make(map[coord]bool)
	q := [][]int{{0, 0}}
	visited[coord{0, 0}] = true
	for len(q) > 0 {
		curr := q[0] // top
		q = q[1:]    // pop
		c, r := curr[0], curr[1]

		// println(r, c, grid[r][c])
		if c == m-1 && r == n-1 {
			return true
		}
		switch grid[r][c] {
		case horizontal:
			if c > 0 && (!visited[coord{r, c - 1}] && (grid[r][c-1] == horizontal || grid[r][c-1] == rightDown || grid[r][c-1] == rightUp)) {
				q = append(q, []int{c - 1, r})
				visited[coord{r, c - 1}] = true
			}
			if c < m-1 && (!visited[coord{r, c + 1}] && (grid[r][c+1] == horizontal || grid[r][c+1] == leftDown || grid[r][c+1] == leftUp)) {
				q = append(q, []int{c + 1, r})
				visited[coord{r, c + 1}] = true
			}
		case vertical:
			if r > 0 && (!visited[coord{r - 1, c}] && (grid[r-1][c] == vertical || grid[r-1][c] == leftDown || grid[r-1][c] == rightDown)) {
				q = append(q, []int{c, r - 1})
				visited[coord{r - 1, c}] = true
			}
			if r < n-1 && (!visited[coord{r + 1, c}] && (grid[r+1][c] == vertical || grid[r+1][c] == leftUp || grid[r+1][c] == rightUp)) {
				q = append(q, []int{c, r + 1})
				visited[coord{r + 1, c}] = true
			}
		case leftDown:
			if c > 0 && (!visited[coord{r, c - 1}] && (grid[r][c-1] == horizontal || grid[r][c-1] == rightDown || grid[r][c-1] == rightUp)) {
				q = append(q, []int{c - 1, r})
				visited[coord{r, c - 1}] = true
			}
			if r < n-1 && (!visited[coord{r + 1, c}] && (grid[r+1][c] == vertical || grid[r+1][c] == leftUp || grid[r+1][c] == rightUp)) {
				q = append(q, []int{c, r + 1})
				visited[coord{r + 1, c}] = true
			}
		case rightDown:
			if c < m-1 && (!visited[coord{r, c + 1}] && (grid[r][c+1] == horizontal || grid[r][c+1] == leftDown || grid[r][c+1] == leftUp)) {
				q = append(q, []int{c + 1, r})
				visited[coord{r, c + 1}] = true
			}
			if r < n-1 && (!visited[coord{r + 1, c}] && (grid[r+1][c] == vertical || grid[r+1][c] == leftUp || grid[r+1][c] == rightUp)) {
				q = append(q, []int{c, r + 1})
				visited[coord{r + 1, c}] = true
			}
		case leftUp:
			if c > 0 && (!visited[coord{r, c - 1}] && (grid[r][c-1] == horizontal || grid[r][c-1] == rightDown || grid[r][c-1] == rightUp)) {
				q = append(q, []int{c - 1, r})
				visited[coord{r, c - 1}] = true
			}
			if r > 0 && (!visited[coord{r - 1, c}] && (grid[r-1][c] == vertical || grid[r-1][c] == leftDown || grid[r-1][c] == rightDown)) {
				q = append(q, []int{c, r - 1})
				visited[coord{r - 1, c}] = true
			}
		case rightUp:
			if c < m-1 && (!visited[coord{r, c + 1}] && (grid[r][c+1] == horizontal || grid[r][c+1] == leftDown || grid[r][c+1] == leftUp)) {
				q = append(q, []int{c + 1, r})
				visited[coord{r, c + 1}] = true
			}
			if r > 0 && (!visited[coord{r - 1, c}] && (grid[r-1][c] == vertical || grid[r-1][c] == leftDown || grid[r-1][c] == rightDown)) {
				q = append(q, []int{c, r - 1})
				visited[coord{r - 1, c}] = true
			}
		}
	}
	return false
}
