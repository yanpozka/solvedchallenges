package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}

func majorityElement(nums []int) []int {
	freqs := map[int]int{}

	for _, n := range nums {
		freqs[n]++
	}

	var res []int
	for n, f := range freqs {
		if f > len(nums)/3 {
			res = append(res, n)
			if len(res) == 3 {
				break
			}
		}
	}
	return res
}
