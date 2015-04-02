package main

import (
	"fmt"
)

const (
	MAXIM    = 4e6
	HALF_MAX = 2e+06
)

func main() {
	var T, N int
	slive := make([]bool, MAXIM)

	for i := 2; i < HALF_MAX; i++ {
		if !slive[i] {
			for k := i + i; k < MAXIM; k += i {
				slive[k] = true // TRUE means NO-prime
			}
		}
	}

	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		fmt.Scanf("%d", &N)
		fmt.Println(slive[N])
	}
}
