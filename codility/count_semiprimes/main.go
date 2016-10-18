package main

import "fmt"

func Solution(N int, P []int, Q []int) []int {
	h := N / 2
	criba := make([]bool, h+1)
	sieve := make([]int, 0, h+1)

	for n := 2; n <= h; n++ {
		if criba[n] == false {
			sieve = append(sieve, n)
			for ix := n + n; ix < len(criba); ix += n {
				criba[ix] = true
			}
		}
	}

	semiprimes := make([]bool, N+1)

	for a, l := 0, len(sieve); a < l; a++ {
		for b := a; b < l; b++ {
			if n := sieve[a] * sieve[b]; n < len(semiprimes) {
				semiprimes[n] = true
			} else {
				break
			}
		}
	}

	acum := make([]int, N+1)
	for ix := 1; ix < len(semiprimes); ix++ {
		acum[ix] = acum[ix-1]
		if semiprimes[ix] {
			acum[ix]++
		}
	}

	var result []int

	for ix, p := range P {
		result = append(result, acum[Q[ix]]-acum[p-1])

		// q := Q[ix]
		// var count int
		// for k := p; k <= q; k++ {
		// 	if semiprimes[k] {
		// 		count++
		// 	}
		// }
		// result = append(result, count)
	}

	return result
}

// var sieve = []int{}

func main() {

	// const Max = 25001
	// criba := make([]bool, Max)
	// // prev := 1

	// for n := 2; n < Max; n++ {
	// 	if criba[n] == false {
	// 		fmt.Print(n, ", ")
	// 		// if n*prev > 50000 {
	// 		// 	 fmt.Println(n, prev)
	// 		// }
	// 		for ix := n + n; ix < Max; ix += n {
	// 			criba[ix] = true
	// 		}
	// 		// prev = n
	// 	}
	// }

	fmt.Println(Solution(26, []int{1, 4, 16, 20}, []int{26, 10, 20, 25}))
}
