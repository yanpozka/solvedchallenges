package main

import (
	"container/heap"
	"fmt"
)

// Given nums1 = [1,7,11], nums2 = [2,4,6],  k = 3
// Return: [1,2],[1,4],[1,6]
// The first 3 pairs are returned from the sequence:
// [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

func main() {
	nums1 := []int{1, 7, 11}
	// nums1 := []int{3, 5, 7, 9}

	// nums2 := []int{}
	// nums2 := []int{2, 4, 6}
	nums2 := []int{2, 4, 9}
	k := 3
	// k := 1

	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}

	pq := new(PriorityQueue)

	for _, a := range nums1 {
		for _, b := range nums2 {
			heap.Push(pq, &Item{a: a, b: b})
		}
	}

	for ; k > 0 && pq.Len() > 0; k-- {
		i := heap.Pop(pq).(*Item)
		result = append(result, []int{i.a, i.b})
	}

	return result
}

type Item struct {
	a, b  int
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].a+pq[i].b < pq[j].a+pq[j].b
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
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
