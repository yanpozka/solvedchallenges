package topkfrequentelements

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums)/2+1)
	for _, n := range nums {
		m[n]++
	}

	pq := make(pQueue, len(m))
	ix := 0
	for n, freq := range m {
		pq[ix] = &itm{
			value: n, priority: freq, index: ix,
		}
		ix++
	}
	heap.Init(&pq)

	res := make([]int, 0, k)
	for pq.Len() > 0 && k > 0 {
		i := heap.Pop(&pq).(*itm)
		res = append(res, i.value)
		k--
	}

	return res
}

type itm struct {
	value, priority int
	index           int
}

type pQueue []*itm

func (pq pQueue) Len() int { return len(pq) }

func (pq pQueue) Less(i, j int) bool {
	if pq[i].priority > pq[j].priority {
		return true
	}
	if pq[i].priority < pq[j].priority {
		return false
	}
	return pq[i].value < pq[j].value
}

func (pq pQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*itm)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *pQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *pQueue) update(item *itm, value, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
