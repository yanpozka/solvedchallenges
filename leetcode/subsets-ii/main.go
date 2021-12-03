package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}

func subsetsWithDup(nums []int) [][]int {
	sets := [][]int{[]int{}}
	hashes := make(map[string]struct{})

	for maskInt, lim := uint(1), uint(math.Pow(2, float64(len(nums)))); maskInt < lim; maskInt++ {
		mask := fmt.Sprintf("%b", maskInt)
		var ns []int

		for ix, c := len(mask)-1, 0; ix >= 0; ix-- {
			if mask[ix] == '1' {
				ns = append(ns, nums[c])
			}
			c++
		}

		sort.Ints(ns)
		var hash strings.Builder
		for _, n := range ns {
			hash.WriteString(strconv.Itoa(n))
		}
		h := hash.String()
		if _, found := hashes[h]; !found {
			sets = append(sets, ns)
			hashes[h] = struct{}{}
		}
	}
	// fmt.Println(hashes)

	return sets
}
