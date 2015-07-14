package main

import (
	"container/list"
	"fmt"
)

func main() {
	var T, N, M, start, x, y int

	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d%d", &N, &M)

		Graph := make([]*list.List, N+1)

		for ix := 1; ix < N+1; ix++ {
			Graph[ix] = list.New()
		}

		for ; M > 0; M-- {
			fmt.Scanf("%d%d", &x, &y)
			Graph[x].PushBack(y)
			Graph[y].PushBack(x)
		}
		fmt.Scanf("%d", &start)

		var dists []int = make([]int, N+1)
		for ix := 1; ix < N+1; ix++ {
			dists[ix] = -1
		}

		dists[start] = 0

		var visited []bool = make([]bool, N+1)
		visited[start] = true

		queue := list.New()
		queue.PushBack(start)

		for queue.Len() > 0 {
			var node int = queue.Remove(queue.Back()).(int)

			for neigh := Graph[node].Front(); neigh != nil; neigh = neigh.Next() {
				adj := neigh.Value.(int)

				if !visited[adj] {
					dists[adj] = dists[node] + 1

					queue.PushBack(adj)
				}
			}
		}

		for ix := 1; ix < N+1; ix++ {
			if ix != start {
				if dists[ix] > 0 {
					fmt.Printf("%d", dists[ix]*6)
				} else {
					fmt.Printf("%d", -1)
				}
				if ix < N {
					fmt.Print(" ")
				}
			}
		}

	} // Test cases
}
