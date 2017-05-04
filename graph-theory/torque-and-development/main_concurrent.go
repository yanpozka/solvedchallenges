package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	var q int
	fmt.Scanf("%d", &q)

	var n, m, cl, cr int
	wg := &sync.WaitGroup{}
	for ; q > 0; q-- {
		fmt.Scanf("%d%d%d%d", &n, &m, &cl, &cr)
		var cg []*list.List
		if cl > cr {
			cg = make([]*list.List, n+1)
			for ix := 1; ix <= n; ix++ {
				cg[ix] = list.New()
			}
		}

		var s, d int
		for ; m > 0; m-- {
			fmt.Scanf("%d%d", &s, &d)

			if cr < cl {
				cg[s].PushBack(d)
				cg[d].PushBack(s)
			}
		}

		if cr > cl {
			fmt.Println(n * cl)
			continue
		}

		wg.Add(1)
		go func(g []*list.List, pn, pm, pcl, pcr int) {
			solution(g, pn, pm, pcl, pcr)
			wg.Done()
		}(cg, n, m, cl, cr)
	}
	wg.Wait()
}

func solution(g []*list.List, n, m, c_lib, c_road int) {

	visited := make([]bool, n+1)
	var result int64

	for ix := 1; ix <= n; ix++ {
		if !visited[ix] { // new component
			no_nodes := bfs(g, ix, 0, visited)

			cost_fix_roads := int64((no_nodes-1)*c_road + c_lib)
			// fmt.Printf("cost_fix_roads: %d\tcost_fix_libs: %d\n", cost_fix_roads, cost_fix_libs)
			// fmt.Println("no_nodes ", no_nodes)
			result += cost_fix_roads
		}
	}

	fmt.Println(result)
}

func bfs(g []*list.List, root, count int, visited []bool) int {
	visited[root] = true

	for e := g[root].Front(); e != nil; e = e.Next() {
		if adj := e.Value.(int); !visited[adj] {
			count += bfs(g, adj, 0, visited)
		}
	}

	return count + 1
}
