package main

import "fmt"

const MaxUint64 = ^uint64(0)
const MaxInt64 = int64(MaxUint64 >> 1)
const MinInt64 = -MaxInt64 - 1

func main() {
	var N, c int
	fmt.Scanf("%d", &N)

	nums := make([][]int, N)
	for i := 0; i < N; i++ {
		nums[i] = make([]int, N)
	}

	fmt.Println(nums)

	for ix := 0; ix < N; ix++ {
		fmt.Scanf("%d", &c)
		// nums = append(nums, c)
	}

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
