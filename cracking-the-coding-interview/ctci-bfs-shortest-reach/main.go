package main

import (
	"container/list"
	"fmt"
)

func main() {
	var q int
	fmt.Scanf("%d", &q)

	for ; q > 0; q-- {
		solution()
	}
}

func solution() {
	var n, m, s int
	fmt.Scanf("%d%d", &n, &m)

	graph := make([]map[int]bool, n+1)
	for ix := 1; ix <= n; ix++ {
		graph[ix] = make(map[int]bool)
	}

	for ; m > 0; m-- {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)

		graph[a][b] = true
		graph[b][a] = true
	}

	fmt.Scanf("%d", &s)

	queue := list.New()
	queue.PushBack(&pair{N: s, L: 0})
	visited := make([]bool, n+1)
	visited[s] = true
	dist := make([]int, n+1)

	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(*pair)
		dist[p.N] = p.L

		for adj := range graph[p.N] {
			if !visited[adj] {
				visited[adj] = true
				queue.PushBack(&pair{N: adj, L: p.L + 1})
			}
		}
	}

	for ix := 1; ix <= n; ix++ {
		if ix != s {
			if dist[ix] == 0 {
				fmt.Printf("%d ", -1)
			} else {
				fmt.Printf("%d ", dist[ix]*6)
			}
		}
	}
	fmt.Println()
}

type pair struct {
	// node and level
	N, L int
}
