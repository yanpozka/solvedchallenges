package main

import (
	"fmt"
	"math"
)

var T, m, M int

func main() {

	base64.NewDecoder()

	fmt.Scanf("%d", &T)

	for ; T > 0; T-- {
		fmt.Scanf("%d%d", &m, &M)
		if m > 6 {
			fmt.Println("YES")
		} else {
			var sum int64 = 0
			var found bool = false

			for i := m; i <= M; i++ {
				sum += int64(math.Pow(10, float64(i)))
				if sum > 1000000 {
					found = true
					break
				}
			}
			if found {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
