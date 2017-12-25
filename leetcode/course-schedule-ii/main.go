// https://play.golang.org/p/qw3teJizyl
// https://leetcode.com/problems/course-schedule-ii/description/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	const n = 2
	fmt.Println(findOrder(n, [][]int{{1, 0}}))
	fmt.Println(findOrder(n, [][]int{{1, 0}, {0, 1}}))

	fmt.Println(findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}})) // true

	fmt.Println(findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 1}})) // false
}

type graph []*list.List

func findOrder(numCourses int, prerequisites [][]int) []int {
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

	res := make([]int, 0, numCourses)
	for set.Len() > 0 {
		n := set.Remove(set.Front()).(int)
		res = append(res, n)

		for e := g[n].Front(); e != nil; {
			m := e.Value.(int)
			inDegree[m]--
			if inDegree[m] < 0 {
				inDegree[m] = 0
			}
			if inDegree[m] == 0 {
				set.PushBack(m)
			}

			next := e.Next()
			g[n].Remove(e)
			e = next
		}
	}

	if isEmptyG(g) {
		return res
	}
	return nil
}

func isEmptyG(g graph) bool {
	for node := range g {
		if g[node].Len() > 0 {
			return false
		}
	}
	return true
}
