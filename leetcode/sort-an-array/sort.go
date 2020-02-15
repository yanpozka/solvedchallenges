package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortArrayHeap([]int{-11}))
	fmt.Println(sortArrayHeap([]int{1, 0}))
	fmt.Println(sortArrayHeap([]int{5, 2, 3, 1}))
	fmt.Println(sortArrayHeap([]int{5, 1, 1, 2, 0, 0}))

	fmt.Println(sortArray([]int{-11}))
	fmt.Println(sortArray([]int{1, 0}))
	fmt.Println(sortArray([]int{5, 2, 3, 1}))
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))

	fmt.Println(sortArrayBurble([]int{-11}))
	fmt.Println(sortArrayBurble([]int{1, 0}))
	fmt.Println(sortArrayBurble([]int{5, 2, 3, 1}))
	fmt.Println(sortArrayBurble([]int{5, 1, 1, 2, 0, 0}))
}

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func sortArrayHeap(nums []int) []int {
	tmp := make([]int, 0, len(nums))
	for _, n := range nums {
		tmp = append(tmp, n)
	}

	ih := intHeap(tmp)
	h := &ih
	heap.Init(h)

	for ix := 0; h.Len() > 0; ix++ {
		nums[ix] = heap.Pop(h).(int)
	}

	return nums
}

func sortArrayBurble(nums []int) []int {
	for range nums {
		var ok bool
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				ok = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		if !ok {
			break
		}
	}

	return nums
}

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
