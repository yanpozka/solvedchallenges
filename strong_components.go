//
// https://play.golang.org/p/V8iCRtiGNr

package main

import (
	"container/list"
	"fmt"
)

type graph []*list.List

func solution(g graph) {
	visited := make([]bool, len(g))
	stack := list.New()

	for ix := 0; ix < len(g); ix++ {
		if !visited[ix] {
			dfs(ix, g, stack, visited)
		}
	}

	tg := inverseGraph(g)
	var comp graph
	visited = make([]bool, len(tg))

	for e := stack.Front(); e != nil; e = e.Next() {
		node := e.Value.(int)

		if !visited[node] {
			l := list.New()
			l.PushBack(node)
			comp = append(comp, l)
			dfs2FindComp(node, tg, comp, visited)
		}
	}

	for u := range comp {
		for e := comp[u].Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value)
		}
		fmt.Println()
	}
}

func dfs2FindComp(r int, g, comp graph, visited []bool) {
	visited[r] = true

	for e := g[r].Front(); e != nil; e = e.Next() {
		node := e.Value.(int)

		if !visited[node] {
			comp[len(comp)-1].PushBack(node)
			dfs2FindComp(node, g, comp, visited)
		}
	}

}

func dfs(r int, g graph, stack *list.List, visited []bool) {
	visited[r] = true

	for e := g[r].Front(); e != nil; e = e.Next() {
		node := e.Value.(int)
		if !visited[node] {
			dfs(node, g, stack, visited)
		}
	}

	stack.PushFront(r)
}

func inverseGraph(g graph) graph {
	new_g := make(graph, len(g))
	for ix := range g {
		new_g[ix] = list.New()
	}

	for u := range g {
		for e := g[u].Front(); e != nil; e = e.Next() {
			v := e.Value.(int)
			new_g[v].PushBack(u)
		}
	}

	return new_g
}

func main() {
	{
		const n = 8

		g := make(graph, n)
		for ix := range g {
			g[ix] = list.New()
		}

		g[0].PushBack(1)
		g[0].PushBack(5)

		g[1].PushBack(2)
		g[1].PushBack(5)

		g[2].PushBack(3)
		g[2].PushBack(6)

		g[4].PushBack(0)

		g[5].PushBack(4)
		g[5].PushBack(6)

		g[6].PushBack(2)

		g[7].PushBack(6)

		// solution(g)
	}

	{
		const n = 8

		g := make(graph, n)
		for ix := range g {
			g[ix] = list.New()
		}

		g[0].PushBack(1)

		g[1].PushBack(2)
		g[1].PushBack(4)
		g[1].PushBack(5)

		g[2].PushBack(3)
		g[2].PushBack(6)

		g[3].PushBack(2)
		g[3].PushBack(7)

		g[4].PushBack(0)
		g[4].PushBack(5)

		g[5].PushBack(6)

		g[6].PushBack(5)
		g[6].PushBack(7)

		g[7].PushBack(7)

		solution(g)
	}
}
