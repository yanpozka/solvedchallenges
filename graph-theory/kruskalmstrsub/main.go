package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//
var tree, rank []int

//
func find(elem int) int {
	if tree[elem] == elem {
		return elem
	}
	return find(tree[elem])
}

//
func merge(a, aP, b, bP int) {
	if rank[aP] > rank[bP] {
		tree[bP] = aP
	} else {
		tree[aP] = bP
		if rank[aP] == rank[bP] {
			rank[bP]++
		}
	}
}

//
func main() {
	var N, M, x, y, w int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	N, _ = strconv.Atoi(parts[0])
	M, _ = strconv.Atoi(parts[1])

	pq := make(PriorityQueue, M*2)

	tree = make([]int, N+1)
	rank = make([]int, N+1)

	for ix := 1; ix < N+1; ix++ { // init forest of trees
		tree[ix] = ix
	}

	var ix int = 0
	for ; M > 0 && scanner.Scan(); M-- {
		parts := strings.Split(scanner.Text(), " ")
		x, _ = strconv.Atoi(parts[0])
		y, _ = strconv.Atoi(parts[1])
		w, _ = strconv.Atoi(parts[2])

		pq[ix] = &Item{value: &edge{src: x, tar: y, weight: w}, priority: w, index: ix}
		ix++
		pq[ix] = &Item{value: &edge{src: y, tar: x, weight: w}, priority: w, index: ix}
		ix++
	}
	scanner.Scan()

	heap.Init(&pq)

	var result, count int

	for pq.Len() > 0 {
		current_edge := heap.Pop(&pq).(*Item).value
		// fmt.Println(current_edge)

		var aP, bP int = find(current_edge.src), find(current_edge.tar)
		if aP != bP { // disjoint sets
			merge(current_edge.src, aP, current_edge.tar, bP)
			result += current_edge.weight
			count++
			if count == N {
				break
			}
		}
	}
	fmt.Println(result)
}

//
type edge struct {
	src, tar, weight int
}

//
type Item struct {
	value           *edge
	priority, index int
}

//
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
