package main

import (
	"container/list"
	"fmt"
)

func main() {
	var q, n, m, cl, cr int
	fmt.Scanf("%d", &q)

	for ; q > 0; q-- {
		fmt.Scanf("%d%d%d%d", &n, &m, &cl, &cr)
		solution(n, m, cl, cr)
	}
}

func solution(n, m, c_lib, c_road int) {
	g := make([]*list.List, n+1)
	for ix := 1; ix <= n; ix++ {
		g[ix] = list.New()
	}

	var s, d int
	for ; m > 0; m-- {
		fmt.Scanf("%d%d", &s, &d)
		g[s].PushBack(d)
		g[d].PushBack(s)
	}

	visited := make([]bool, n+1)
	for ix := 1; ix <= n; ix++ {
		if !visited[ix] { // new graph
			fmt.Println("--------------")
			bfs(g, ix, visited)
			fmt.Println("--------------")
		}
	}
}

func bfs(g []*list.List, root int, visited []bool) {
	visited[root] = true
	fmt.Println(root)

	for e := g[root].Front(); e != nil; e = e.Next() {
		if adj := e.Value.(int); !visited[adj] {
			bfs(g, adj, visited)
		}
	}
}
