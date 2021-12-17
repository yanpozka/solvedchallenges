package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	freqs := map[int]int{}

	for _, n := range nums {
		freqs[n]++
	}

	var res int
	for n, f := range freqs {
		if f > len(nums)/2 {
			res = n
			break
		}
	}
	return res
}
