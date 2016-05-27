package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

const MaxUint = ^uint(0)
const INF = int(MaxUint >> 1)

var Graph []*list.List

func main() {
	var N, M, start, end, x, y, w int
	fmt.Scan(&N, &M)

	Graph = make([]*list.List, N+1)

	for ix := 1; ix < N+1; ix++ {
		Graph[ix] = list.New()
	}

	for ; M > 0; M-- {
		fmt.Scan(&x, &y, &w)

		Graph[x].PushBack(&edge{index: y, weight: w})
		Graph[y].PushBack(&edge{index: x, weight: w})
	}
	fmt.Scan(&start, &end)

	dists, prev := dijkstra(start, N) // !!

	if dists[end] != INF {
		result := []int{end}
		a := end
		c := prev[a]
		for ; c != start && c != 0; c = prev[a] {
			result = append(result, c)
			a = c
		}
		result = append(result, c)
		sum_or := uint32(0)

		for l := len(result) - 1; l >= 0; l-- {
			if l-1 < 0 {
				break
			}
			a, b := result[l], result[l-1]

			min := INF
			for elem := Graph[a].Front(); elem != nil; elem = elem.Next() {
				if e, isType := elem.Value.(*edge); isType && e.index == b {
					if e.weight < min {
						min = e.weight
					}
					// fmt.Println(min)
				}
			}
			if min == INF {
				panic("min inf")
			}
			// fmt.Println()
			sum_or |= uint32(min)
		}
		// fmt.Println(dists[end])
		fmt.Println(sum_or)
	} else {
		fmt.Println(-1)
	}
}

func dijkstra(start, N int) ([]int, []int) {
	pq := make(PriorityQueue, Graph[start].Len())
	dists := make([]int, N+1)
	prev := make([]int, N+1)

	prev[start] = start

	for ix := 1; ix < N+1; ix++ {
		dists[ix] = INF
	}
	var i int = 0
	for e := Graph[start].Front(); e != nil; e = e.Next() {
		c := e.Value.(*edge)
		other := &edge{index: c.index, weight: c.weight}
		pq[i] = &Item{value: other, priority: other.weight, index: i}
		i++ // !!

		if c.weight < dists[c.index] {
			dists[other.index] = other.weight
			prev[other.index] = start
		}
	}

	heap.Init(&pq)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).value

		for adj := Graph[current.index].Front(); adj != nil; adj = adj.Next() {
			u := adj.Value.(*edge)
			if u.index == start {
				continue
			}

			if dists[current.index] != INF && dists[current.index]+u.weight < dists[u.index] {
				dists[u.index] = dists[current.index] + u.weight

				im := &Item{
					value:    &edge{index: u.index, weight: dists[u.index]},
					priority: dists[u.index],
				}
				heap.Push(&pq, im)
				pq.update(im, im.value, dists[u.index])

				prev[u.index] = current.index // prev[v] ← u
			}
		}
	}
	return dists, prev
}

//
type edge struct {
	index, weight int
}

//
type Item struct {
	value           *edge
	priority, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value *edge, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
