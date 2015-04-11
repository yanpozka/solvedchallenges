package main

import (
	"fmt"
	// "math"
)

func main() {
	var T, N, count int64
	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d", &N)
		var first, second int64 = 1, 2
		count = 0

		for first < N {
			if first%2 == 0 {
				count += first
			}
			first, second = second, first+second
		}
		fmt.Println(count)
	}
}
