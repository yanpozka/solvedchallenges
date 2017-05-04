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

	if c_road > c_lib {
		fmt.Println(n * c_lib)
		return
	}

	visited := make([]bool, n+1)
	var result int64

	for ix := 1; ix <= n; ix++ {
		if !visited[ix] { // new component
			no_nodes := dfs(g, ix, 0, visited)
			cost_fix_roads := int64((no_nodes-1)*c_road + c_lib)
			result += cost_fix_roads
		}
	}

	fmt.Println(result)
}

func dfs(g []*list.List, root, count int, visited []bool) int {
	visited[root] = true

	for e := g[root].Front(); e != nil; e = e.Next() {
		if adj := e.Value.(int); !visited[adj] {
			count += dfs(g, adj, 0, visited)
		}
	}

	return count + 1
}
