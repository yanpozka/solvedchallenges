package main

import "fmt"

func main() {
	var N, T int
	fmt.Scanf("%d %d", &N, &T)

	nums := make([]int, N)

	for ix := 0; ix < N; ix++ {
		fmt.Scan(&nums[ix])
	}

	var s, end int
	for ; T > 0; T-- {
		fmt.Scanf("%d %d", &s, &end)
		min := nums[s]
		for s = s + 1; s <= end; s++ {
			if c := nums[s]; c < min {
				min = c
			}
		}
		fmt.Println(min)
	}
}
