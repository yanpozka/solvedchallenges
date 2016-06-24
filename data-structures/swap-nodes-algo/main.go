package main

import (
	"container/list"
	"fmt"
)

const (
	LEFT = iota
	RIGTH
)

type btree [][2]int

func (t btree) inorder(root int) {
	if root == -1 {
		return
	}
	t.inorder(t[root][LEFT])
	fmt.Print(root, " ")
	t.inorder(t[root][RIGTH])
}

//
type node struct {
	index, h int
}

func main() {
	var N, T, a, b int
	fmt.Scan(&N)
	t := make(btree, N+1)

	for ix := 1; ix <= N; ix++ {
		fmt.Scan(&a, &b)
		t[ix][LEFT], t[ix][RIGTH] = a, b
	}

	fmt.Scan(&T)
	for ; T > 0; T-- {
		fmt.Scan(&a)

		q := list.New()
		q.PushBack(&node{index: 1, h: 1})

		for q.Len() > 0 {
			current := q.Remove(q.Front()).(*node)

			if current.h >= a && current.h%a == 0 {
				root := current.index
				t[root][LEFT], t[root][RIGTH] = t[root][RIGTH], t[root][LEFT]
			}

			if left := t[current.index][LEFT]; left != -1 {
				q.PushBack(&node{index: left, h: current.h + 1})
			}
			if right := t[current.index][RIGTH]; right != -1 {
				q.PushBack(&node{index: right, h: current.h + 1})
			}
		}

		t.inorder(1)
		fmt.Println()
	}
}
