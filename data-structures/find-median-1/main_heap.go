package main

import (
	"container/heap"
	"fmt"
)

func main() {
	var t, n int
	fmt.Scan(&t)

	dec := &DecIntHeap{&IntHeap{}}
	h := &IntHeap{}

	for ; t > 0; t-- {
		fmt.Scan(&n)

		// add
		if dec.Len() == 0 || (*dec.IntHeap)[0] > n {
			heap.Push(dec, n)
		} else {
			heap.Push(h, n)
		}

		// rebalance
		diff := h.Len() - dec.Len()
		if diff >= 2 {
			if h.Len() > 0 {
				heap.Push(dec, heap.Pop(h))
			}
		} else if diff <= -2 {
			if dec.Len() > 0 {
				heap.Push(h, heap.Pop(dec))
			}
		}

		if total := uint32(h.Len() + dec.Len()); total&1 == 1 { // even
			if h.Len() > dec.Len() {
				fmt.Printf("%.1f\n", float32((*h)[0]))
			} else {
				fmt.Printf("%.1f\n", float32((*dec.IntHeap)[0]))
			}
		} else { // odd
			fmt.Printf("%.1f\n", float64((*h)[0]+(*dec.IntHeap)[0])/2.0)
		}
	}
}

// min-heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// max-heap
type DecIntHeap struct {
	*IntHeap
}

func (ih *DecIntHeap) Less(i, j int) bool { return !ih.IntHeap.Less(i, j) }
