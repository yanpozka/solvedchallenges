// https://play.golang.org/p/pRUVF14YwM5
//
// https://en.wikipedia.org/wiki/Topological_sorting
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
}

type graph []*list.List

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make(graph, numCourses)
	for ix := range g {
		g[ix] = list.New()
	}
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		src, dst := p[1], p[0]
		g[src].PushBack(dst)
		inDegree[dst]++
	}

	set := list.New()
	for node, degree := range inDegree {
		if degree == 0 {
			set.PushBack(node)
		}
	}

	for set.Len() > 0 {
		n := set.Remove(set.Front()).(int)

		for e := g[n].Front(); e != nil; {
			m := e.Value.(int)
			inDegree[m]--
			if inDegree[m] == 0 {
				set.PushBack(m)
			}

			next := e.Next()
			g[n].Remove(e)
			e = next
		}
	}

	return isEmptyG(g)
}

func isEmptyG(g graph) bool {
	for node := range g {
		if g[node].Len() > 0 {
			return false
		}
	}
	return true
}
