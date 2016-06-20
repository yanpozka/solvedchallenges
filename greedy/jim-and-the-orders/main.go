package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	var N, t, d int
	fmt.Scan(&N)

	nums := make([]int, N)
	m := make(map[int]*list.List)

	for ix := 0; ix < N; ix++ {
		fmt.Scan(&t, &d)
		key := t + d
		nums[ix] = key

		if _, contains := m[key]; !contains {
			m[key] = list.New()
		}

		if m[key].Len() == 0 {
			m[key].PushBack(ix)
		} else {
			e := m[key].Back()
			for ; e != m[key].Front() && e.Value.(int) > ix; e = e.Prev() {
			}

			m[key].InsertAfter(ix, e) // original_index and new index (ix)
		}
	}

	sort.Ints(nums)
	prev := -1

	for _, n := range nums {
		if n <= prev {
			continue
		}
		prev = n

		for e := m[n].Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value.(int)+1, " ")
		}
	}
	fmt.Println()
}
