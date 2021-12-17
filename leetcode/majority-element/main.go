package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) int {
	freqs := map[int]int{}
	var res int
	for _, n := range nums {
		freqs[n]++
		if freqs[n] > len(nums)/2 {
			res = n
			break
		}
	}
	return res
}
