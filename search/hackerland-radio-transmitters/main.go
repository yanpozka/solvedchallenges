package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, a int
	fmt.Scanf("%d%d", &n, &k)
	nums := make([]int, n)

	for ix := range nums {
		fmt.Scanf("%d", &a)
		nums[ix] = a
	}
	sort.Ints(nums)

	var count int

	for ix := 0; ix < n; ix++ {
		fmt.Println()
		fmt.Println("ix", ix)
		v := nums[ix]
		centerIndex := -1

		for p_c := (k + ix) % n; p_c > ix; p_c-- {
			if nums[p_c] <= k+v {
				centerIndex = p_c
				break
			}
		}
		fmt.Println("centerIndex", centerIndex)

		count++
		if centerIndex == -1 {
			continue
		}

		ix = centerIndex // !!

		for a := centerIndex + 1; a <= centerIndex+k && a < n; a++ {
			if nums[a] <= nums[centerIndex]+k {
				fmt.Printf("%d <= %d \n", nums[a], nums[centerIndex]+k)
				ix++
			} else {
				break
			}
		}
	}

	fmt.Println(count)
}
