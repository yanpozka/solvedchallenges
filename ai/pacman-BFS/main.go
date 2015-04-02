package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

const (
	EMPTY = 45 // -
	WALL  = 37 // %
)

type Pair struct {
	R, C, IxParent int
}

var pac_row, pac_col, food_row, food_col, N, M int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Scanf("%d%d", &pac_row, &pac_col)
	fmt.Scanf("%d%d", &food_row, &food_col)
	fmt.Scanf("%d%d", &N, &M)

	Matrix := make([][]uint8, N)

	for i := 0; i < N && scanner.Scan(); i++ {
		var line string = scanner.Text()
		Matrix[i] = make([]uint8, M)

		for k := range line {
			Matrix[i][k] = line[k]
		}
	}
	Matrix[pac_row][pac_col] = WALL // pacman position

	stack_result := make([]*Pair, 0, 100)
	queue_coords := list.New()
	queue_coords.PushFront(&Pair{R: pac_row, C: pac_col, IxParent: -1})

	for queue_coords.Len() > 0 {
		elem := queue_coords.Front()
		queue_coords.Remove(elem)

		p := elem.Value.(*Pair)
		stack_result = append(stack_result, p)

		var r, c int = p.R, p.C
		if r == food_row && c == food_col {
			break // found it
		}
		var indexP int = len(stack_result) - 1

		if Matrix[r-1][c] != WALL { // UP
			queue_coords.PushBack(&Pair{R: r - 1, C: c, IxParent: indexP})
			Matrix[r-1][c] = WALL
		}
		if Matrix[r][c-1] != WALL { // LEFT
			queue_coords.PushBack(&Pair{R: r, C: c - 1, IxParent: indexP})
			Matrix[r][c-1] = WALL
		}
		if Matrix[r][c+1] != WALL { // RIGHT
			queue_coords.PushBack(&Pair{R: r, C: c + 1, IxParent: indexP})
			Matrix[r][c+1] = WALL
		}
		if Matrix[r+1][c] != WALL { // DOWN
			queue_coords.PushBack(&Pair{R: r + 1, C: c, IxParent: indexP})
			Matrix[r+1][c] = WALL
		}
	}

	var len_result int = len(stack_result)
	fmt.Println(len_result)
	for ix := range stack_result {
		fmt.Println(stack_result[ix].R, stack_result[ix].C)
	}

	result_list := list.New()
	result_list.PushFront(stack_result[len_result-1])

	for ix_parent := stack_result[len_result-1].IxParent; ix_parent != -1; ix_parent = stack_result[ix_parent].IxParent {

		result_list.PushFront(stack_result[ix_parent])
	}

	fmt.Println(result_list.Len() - 1)
	for node := result_list.Front(); node != nil; node = node.Next() {
		p := node.Value.(*Pair)
		fmt.Println(p.R, p.C)
	}
}
