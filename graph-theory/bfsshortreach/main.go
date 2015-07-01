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
	var T, N, M, start, x, y int

	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d%d", &N, &M)

		Graph = make([]*list.List, N+1)

		for ix := 1; ix < N+1; ix++ {
			Graph[ix] = list.New()
		}

		for ; M > 0; M-- {
			fmt.Scanf("%d%d", &x, &y)
			Graph[x].PushBack(y)
			Graph[y].PushBack(x)
		}
		fmt.Scanf("%d", &start)

		dists := dijkstra(start, N)

		for ix := 1; ix < N+1; ix++ {
			if ix != start {
				if dists[ix] != INF {
					fmt.Printf("%d", dists[ix]*6)
				} else {
					fmt.Printf("%d", -1)
				}
				if ix < N {
					fmt.Print(" ")
				}
			}
		}
		if T > 1 {
			fmt.Println()
		}
	} // Test cases
}

func dijkstra(start, N int) []int {
	pq := make(PriorityQueue, Graph[start].Len())
	var dists []int = make([]int, N+1)

	for ix := 1; ix < N+1; ix++ {
		dists[ix] = INF
	}
	var i, c = 0, 0
	for e := Graph[start].Front(); e != nil; e = e.Next() {
		c = e.Value.(int)
		pq[i] = &Item{value: c, priority: 1, index: i}
		i++
		dists[c] = 1
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)

		for adj := Graph[item.value].Front(); adj != nil; adj = adj.Next() {
			u := adj.Value.(int)
			if u == start {
				continue
			}
			if dists[item.value] != INF && dists[item.value]+1 < dists[u] /* [item,adj] == 1 */ {
				dists[u] = dists[item.value] + 1

				im := &Item{
					value:    u,
					priority: dists[u],
				}
				heap.Push(&pq, im)
				pq.update(im, im.value, dists[u])
			}
		}
	}
	return dists
}

type Item struct {
	value, priority, index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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

func (pq *PriorityQueue) update(item *Item, value, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
