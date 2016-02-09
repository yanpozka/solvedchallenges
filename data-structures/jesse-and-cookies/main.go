package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	K, _ := strconv.ParseInt(scanner.Text(), 10, 0)

	h := &IntHeap{}
	for ix := 0; ix < N && scanner.Scan(); ix++ {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		heap.Push(h, n)
	}
	var count int
	var ok bool

	for h.Len() > 1 {
		f := heap.Pop(h).(int64)
		if f >= K {
			ok = true
			break
		}
		s := heap.Pop(h).(int64)
		heap.Push(h, f+(2*s))
		count++
	}

	if ok {
		fmt.Println(count)
	} else {
		f := heap.Pop(h).(int64)
		if f >= K {
			fmt.Println(count)
		} else {
			fmt.Println(-1)
		}
	}
}

//
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
