package main

import "fmt"

const MaxUint64 = ^uint64(0)
const MaxInt64 = int64(MaxUint64 >> 1)
const MinInt64 = -MaxInt64 - 1

func main() {
	var N, c int
	fmt.Scanf("%d", &N)

	nums := make([]int, 0, N)

	var max, csum int64 = MinInt64, 0

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d", &c)

		for _, val := range nums {
			csum = int64(val * c)
			if csum > max {
				max = csum
			}
		}
		// O ( n ^ 2 )
		nums = append(nums, c)
	}

	for rang, half := 2, N/2; rang <= half; rang++ {
		// fmt.Println("Range=", rang)
		for ix := 0; ix < rang; ix++ {
			var start, end int = ix, ix + rang
			// fmt.Println("ix=", start, end)
			csum = 0

			for k := end; k <= N-rang; k++ {
				csum = mult(nums[start:end], nums[end:end+rang], rang)
				if csum > max {
					// fmt.Println(start, end, end, end+rang)
					// fmt.Println(csum)
					max = csum
				}
			}
		}
	}
	fmt.Println(max)
}

func mult(a []int, b []int, lenght int) int64 {
	var i int = 0
	var sum int64 = 0
	for k := lenght - 1; k >= 0 && i < lenght; k-- {
		sum += int64(a[i] * b[k])
		i++
	}
	return sum
}
