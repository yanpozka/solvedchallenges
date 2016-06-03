package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var N, action, val int
	fmt.Scan(&N)
	h := &intHeap{}

	for ; N > 0; N-- {
		fmt.Scan(&action)

		switch action {
		case 1:
			fmt.Scan(&val)
			heap.Push(h, val)

		case 2:
			fmt.Scan(&val)
			var i int
			for ix, l := 0, h.Len(); ix < l; ix++ {
				if (*h)[ix] == val {
					i = ix
					break
				}
			}
			heap.Remove(h, i)

		case 3:
			fmt.Println((*h)[0])
		}
	}
}

//
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
