package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K, c int

	fmt.Scanf("%d%d", &N, &K)
	// fmt.Println(N, K)
	var boys []int = make([]int, N)
	var girls []int = make([]int, N)

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d", &c)
		boys[ix] = c
	}

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d", &c)
		girls[ix] = c
	}

	sort.Sort(sort.IntSlice(boys))
	sort.Sort(sort.IntSlice(girls))

	var count int = 0
	for b, g := 0, 0; b < N && g < N; g++ {
		if abs(boys[b]-girls[g]) <= K {
			b++
			count++
		}
	}
	fmt.Println(count)
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
