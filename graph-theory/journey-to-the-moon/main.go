package main

import (
	"container/list"
	"fmt"
)

var (
	graph   []*list.List
	visited []bool
)

func bfs(root int) int {
	visited[root] = true
	var count int

	for e := graph[root].Front(); e != nil; e = e.Next() {
		if adj := e.Value.(int); visited[adj] == false {
			count += bfs(adj)
		}
	}
	return count + 1
}

func main() {
	var N, P, a, b int
	fmt.Scan(&N, &P)

	graph = make([]*list.List, N)
	visited = make([]bool, N)

	for ix := 0; ix < N; ix++ {
		graph[ix] = list.New()
	}

	for ; P > 0; P-- {
		fmt.Scan(&a, &b)
		graph[a].PushBack(b)
		graph[b].PushBack(a)
	}

	gs := make([]uint64, 0, N+1)

	for ix := 0; ix < N; ix++ {
		if visited[ix] == false {
			gs = append(gs, uint64(bfs(ix)))
		}
	}

	sum := uint64((N * (N - 1)) / 2)
	var ways uint64

	for ix, S := 0, len(gs); ix < S; ix++ {
		ways += (gs[ix] * (gs[ix] - 1) / 2)
	}

	fmt.Println(sum - ways)
}
