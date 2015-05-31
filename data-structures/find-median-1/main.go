package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, numb int

	fmt.Scanf("%d", &N)

	var numbers []int = make([]int, 0, N)

	for ; N > 0; N-- {
		fmt.Scanf("%d", &numb)
		numbers = insert(numbers, numb)

		var ln int = len(numbers)
		var r float32 = 0

		if ln%2 == 0 {
			r = float32(numbers[ln/2]+numbers[ln/2-1]) / 2.0
		} else {
			r = float32(numbers[ln/2])
		}
		fmt.Printf("%.1f\n", r)
	}

	fmt.Println(numbers)
}

func insert(nms []int, n int) []int {
	var ln int = len(nms)
	if ln == 0 {
		return append(nms, n)
	}

	var index int = sort.SearchInts(nms, n)
	// fmt.Println("found=", index)

	// return append(nms[:index], append([]int{n}, nms[index:]...)...)
	nms = append(nms, 0)
	copy(nms[index+1:], nms[index:])
	nms[index] = n

	return nms
}
