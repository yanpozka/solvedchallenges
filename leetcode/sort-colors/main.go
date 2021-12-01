package main

import "fmt"

func main() {
	{
		arr := []int{2, 0, 2, 1, 1, 0}
		sortColors(arr)
		fmt.Println(arr)
		arr = []int{2, 0, 1}
		sortColors(arr)
		fmt.Println(arr)
	}
}

func sortColors(nums []int) {
	var counts [3]int
	for _, n := range nums {
		counts[n]++
	}
	var ix int
	for color, c := range counts {
		for ; c > 0; c-- {
			nums[ix] = color
			ix++
		}
	}
}
