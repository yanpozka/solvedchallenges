package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sr := bufio.NewScanner(os.Stdin)
	sr.Split(bufio.ScanWords)

	sr.Scan()
	N, _ := strconv.Atoi(sr.Text())

	sr.Scan()
	K, _ := strconv.Atoi(sr.Text())

	h := &IntHeap{}

	for ix := 0; ix < N && sr.Scan(); ix++ {
		n, _ := strconv.Atoi(sr.Text())
		heap.Push(h, n)
	}

	var count int
	for {
		c := heap.Pop(h).(int)
		if K >= c {
			count++
			K -= c
		} else {
			break
		}
	}

	fmt.Println(count)
}

//
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
