package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type pair struct {
	r, c int
	wave bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	mrune := "M"[0]
	space := "."[0]
	TREE := "X"[0]
	START := "*"[0]

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0; T-- {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		M, _ := strconv.Atoi(scanner.Text())

		var matrix = make([][]byte, N)

		var m_row, m_col int = -1, -1
		for ix := 0; ix < N && scanner.Scan(); ix++ {
			var line = scanner.Bytes()

			if m_row == -1 {
				for k := 0; k < M; k++ {
					if line[k] == mrune {
						m_row, m_col = ix, k
						break
					}
				}
			}
			matrix[ix] = line
		}
		scanner.Scan()
		K, _ := strconv.Atoi(scanner.Text())

		queue := list.New()
		root := &pair{r: m_row, c: m_col}
		queue.PushBack(root)

		parents := make(map[*pair]*pair)
		parents[root] = nil

		var count, adjs int = 0, 0
		var last *pair

		for queue.Len() > 0 {
			c := queue.Remove(queue.Front()).(*pair)
			if matrix[c.r][c.c] == START {
				last = c
				break
			}
			adjs = 0
			matrix[c.r][c.c] = TREE

			if c.c-1 >= 0 && (matrix[c.r][c.c-1] == space || matrix[c.r][c.c-1] == START) { // LEFT
				p := &pair{r: c.r, c: c.c - 1}
				queue.PushBack(p)
				parents[p] = c
				adjs++
			}
			if c.c+1 < M && (matrix[c.r][c.c+1] == space || matrix[c.r][c.c+1] == START) { // RIGHT
				p := &pair{r: c.r, c: c.c + 1}
				queue.PushBack(p)
				parents[p] = c
				adjs++
			}
			if c.r-1 >= 0 && (matrix[c.r-1][c.c] == space || matrix[c.r-1][c.c] == START) { // UP
				p := &pair{r: c.r - 1, c: c.c}
				queue.PushBack(p)
				parents[p] = c
				adjs++
			}
			if c.r+1 < N && (matrix[c.r+1][c.c] == space || matrix[c.r+1][c.c] == START) { // DOWN
				p := &pair{r: c.r + 1, c: c.c}
				queue.PushBack(p)
				parents[p] = c
				adjs++
			}

			if adjs > 1 {
				c.wave = true
			}
		}

		for c := last; c != nil; c = parents[c] {
			if c.wave {
				count++
				if count > K {
					break
				}
			}
		}

		if count == K {
			fmt.Println("Impressed")
		} else {
			fmt.Println("Oops!")
		}
	}
}
