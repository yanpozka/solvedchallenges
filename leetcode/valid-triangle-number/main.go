package main

import (
	"fmt"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
}

func triangleNumber(nums []int) int {
	var maxNum int
	var counts [1001]int
	for _, n := range nums {
		counts[n]++
		if n > maxNum {
			maxNum = n
		}
	}

	var orNums []int
	for n, c := range counts[:maxNum+1] {
		for ; c > 0; c-- {
			orNums = append(orNums, n)
		}
	}

	var ts int
	for ix := 2; ix < len(orNums); ix++ {
		a, b := 0, ix-1
		for a < b {
			if sum := orNums[a] + orNums[b]; sum > orNums[ix] {
				ts += b - a
				b--
			} else {
				a++
			}
		}
	}
	return ts
}
