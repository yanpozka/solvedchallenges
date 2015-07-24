package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const INF = int(MaxUint >> 1)

var Graph []*list.List

func main() {
	var T, N, M, start, x, y, w int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ = strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		parts := strings.Split(scanner.Text(), " ")

		N, _ = strconv.Atoi(parts[0])
		M, _ = strconv.Atoi(parts[1])

		Graph = make([]*list.List, N+1)

		for ix := 1; ix < N+1; ix++ {
			Graph[ix] = list.New()
		}

		for ; M > 0 && scanner.Scan(); M-- {
			parts := strings.Split(scanner.Text(), " ")
			x, _ = strconv.Atoi(parts[0])
			y, _ = strconv.Atoi(parts[1])
			w, _ = strconv.Atoi(parts[2])

			Graph[x].PushBack(&edge{index: y, weight: w})
			Graph[y].PushBack(&edge{index: x, weight: w})
		}
		scanner.Scan()
		start, _ = strconv.Atoi(scanner.Text())

		dists := dijkstra(start, N) // !!

		for ix := 1; ix < N+1; ix++ {
			if ix != start {
				if dists[ix] != INF {
					fmt.Printf("%d ", dists[ix])
				} else {
					fmt.Printf("%d ", -1)
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
	var i int = 0
	for e := Graph[start].Front(); e != nil; e = e.Next() {
		c := e.Value.(*edge)
		other := &edge{index: c.index, weight: c.weight}
		pq[i] = &Item{value: other, priority: other.weight, index: i}
		i++ // !!

		if c.weight < dists[c.index] {
			dists[other.index] = other.weight
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
			}
		}
	}
	return dists
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
