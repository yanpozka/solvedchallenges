// https://play.golang.org/p/qOX4-NjYLX
//
package main

import (
	"container/list"
	"fmt"
)

func main() {
	const n = 2
	fmt.Println(canFinish(n, [][]int{{1, 0}}))
	fmt.Println(canFinish(n, [][]int{{1, 0}, {0, 1}}))

	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}})) // true

	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 1}})) // false

	//
	fmt.Println(canFinish(8, [][]int{
		{1, 0},
		{2, 6},
		{1, 7},
		{5, 1},
		{6, 4},
		{7, 0},
		{0, 5},
	})) // false
}

type graph []*list.List

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make(graph, numCourses)
	for ix := range g {
		g[ix] = list.New()
	}

	for _, p := range prerequisites {
		src, dst := p[1], p[0]
		g[src].PushBack(dst)
	}

	visited := make([]bool, numCourses)
	rec := make([]bool, numCourses)

	for node := 0; node < numCourses; node++ {
		if visited[node] {
			continue
		}
		if hasCycle(g, node, visited, rec) {
			return false
		}
	}

	return true
}

func hasCycle(g graph, root int, checked, rec []bool) bool {
	checked[root] = true
	rec[root] = true

	for e := g[root].Front(); e != nil; e = e.Next() {
		adj := e.Value.(int)
		if !checked[adj] {
			if hasCycle(g, adj, checked, rec) {
				return true
			}
		} else if rec[adj] {
			return true
		}
	}
	rec[root] = false
	return false
}
