package main

import (
	"fmt"
	// "math"
)

const (
	MAXIM    = 4e6
	HALF_MAX = 2e+06
)

func main() {
	var T, N int
	slive := make([]bool, MAXIM)
	primes := make([]int, 0, 300)

	for i := 2; i < HALF_MAX; i++ {
		if !slive[i] { // i is prime
			primes = append(primes, i)
			for k := i + i; k < MAXIM; k += i {
				slive[k] = true // TRUE means NO-prime
			}
		}
	}

	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		fmt.Scanf("%d", &N)

		var r int = int(math.Ceil(math.Sqrt(float64(N))))
		var found bool = false
		for ; r > 1; r-- {
			if !slive[r] && N%r == 0 {
				fmt.Println(r)
				found = true
				break
			}
		}

		if !found {
			fmt.Println(N)
		}
	}
}
