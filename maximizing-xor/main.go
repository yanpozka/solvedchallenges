package main

import "fmt"

func main() {
	var L, R, max, a, b uint16

	fmt.Scanf("%d\n%d", &L, &R)
	max = 0

	for a = L; a <= R; a++ {
		for b = a + 1; b <= R; b++ {
			temp := a ^ b
			if temp > max {
				max = temp
			}
		}
	}

	fmt.Println(max)
}
