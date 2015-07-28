package main

import "fmt"

func main() {
	var N int

	fmt.Scanf("%d", &N)

	for row := 0; row < N; row++ {

		for c := 0; c < N-2-row+1; c++ {
			fmt.Print(" ")
		}
		for c := N - 2 - row + 1; c < N; c++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}
