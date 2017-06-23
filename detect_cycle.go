package main

import (
	"container/list"
	"fmt"
)

func detectCycles(n int, g graph) {
	visited := make([]bool, n)
	parents := make([]int, n)
	for ix := range parents {
		parents[ix] = -1
	}

	c := make(map[edge]struct{})

	var count int
	for ix := 0; ix < n; ix++ {
		if !visited[ix] {
			dfs(ix, g, parents, visited, c)
			count += len(c)
		}
	}

	fmt.Printf("number of cycles: %d\n", count)
}

type edge struct {
	src, dst int
}

func dfs(r int, g graph, parents []int, visited []bool, c map[edge]struct{}) {
	visited[r] = true

	for e := g[r].Front(); e != nil; e = e.Next() {
		adj := e.Value.(int)

		if visited[adj] { // new cycle
			if parents[r] != adj {
				a, b := edge{r, adj}, edge{adj, r}
				_, containsA := c[a]
				if _, containsB := c[b]; !containsA && !containsB {
					c[a] = struct{}{}
				}
			}
		} else {
			parents[adj] = r
			dfs(adj, g, parents, visited, c)
		}
	}
}

/*

0 -- 1 -- 3
 \   |  /
  \  | /
     2

*/

func main() {
	const n = 4

	g := make(graph, n)
	for ix := range g {
		g[ix] = list.New()
	}

	g[0].PushBack(1)
	g[0].PushBack(2)

	g[1].PushBack(0)
	g[1].PushBack(2)
	g[1].PushBack(3)

	g[2].PushBack(0)
	g[2].PushBack(1)
	g[2].PushBack(3)

	g[3].PushBack(2)
	g[3].PushBack(1)

	detectCycles(n, g)
}

type graph []*list.List
