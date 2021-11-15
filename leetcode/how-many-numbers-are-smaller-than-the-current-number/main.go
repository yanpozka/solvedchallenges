package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var freqs [101]int
	var acc [101]int

	var max int
	for _, n := range nums {
		freqs[n]++
		if n > max {
			max = n
		}
	}
	var count int
	for n, f := range freqs[:max+1] {
		acc[n] = count
		count += f
	}

	sol := make([]int, len(nums))
	for ix, n := range nums {
		sol[ix] = acc[n]
	}
	return sol
}
