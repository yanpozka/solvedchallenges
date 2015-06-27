package main

import "fmt"

func main() {
	var N, P, X, current, max, ix_max int
	fmt.Scanf("%d%d%d", &N, &P, &X)

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d", &current)

		val := current * P
		if val > max {
			max = val
			ix_max = ix + 1
		}
		P -= X
	}
	fmt.Println(ix_max)
}
