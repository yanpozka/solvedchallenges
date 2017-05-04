package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func main() {
	var q, n, m, cl, cr int
	scanner = bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &q)

	for ; q > 0; q-- {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d%d%d%d", &n, &m, &cl, &cr)

		solution(n, m, cl, cr)
	}
}

func solution(n, m, c_lib, c_road int) {
	var g []*list.List
	if c_road > c_lib == false {
		g = make([]*list.List, n+1)
		for ix := 1; ix <= n; ix++ {
			g[ix] = list.New()
		}
	}

	var s, d int
	for ; m > 0; m-- {
		scanner.Scan()
		bb := scanner.Bytes()

		if c_road > c_lib == false {
			s, d = 0, 0
			var ix int
			for _, v := range bb {
				ix++
				if v == ' ' {
					break
				}
				s = s*10 + int(v-'0')
			}

			for ; ix < len(bb); ix++ {
				d = d*10 + int(bb[ix]-'0')
			}

			g[s].PushBack(d)
			g[d].PushBack(s)
		}
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
