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

var Graph []*list.List

func main() {
	var N, M, start, x, y, w int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
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

	var visited []bool = make([]bool, N+1)

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{value: &edge{index: start, weight: 0}, priority: 0, index: 0}
	heap.Init(&pq)

	var result int

	for pq.Len() > 0 {
		current_edge := heap.Pop(&pq).(*Item).value

		if visited[current_edge.index] {
			continue
		}
		visited[current_edge.index] = true

		result += current_edge.weight // !!

		for adj := Graph[current_edge.index].Front(); adj != nil; adj = adj.Next() {
			u := adj.Value.(*edge)
			if visited[u.index] {
				continue
			}

			item := &Item{value: u, priority: u.weight}
			heap.Push(&pq, item)
			pq.update(item, item.value, u.weight)
		}
	}

	fmt.Println(result)
}

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
