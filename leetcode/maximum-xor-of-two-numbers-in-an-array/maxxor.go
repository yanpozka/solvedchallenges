package main

import "fmt"

func main() {
	in := []int{2, 3, 5, 8, 10, 25}
	fmt.Println(findMaximumXOR(in))
}

func findMaximumXOR(nums []int) int {
	var bits [32][]int
	var maxBit int

	for _, n := range nums {
		c := countBits(n)
		bits[c] = append(bits[c], n)
		if c > maxBit {
			maxBit = c
		}
	}
	fmt.Println(bits[:maxBit+1])

	return bits[maxBit][0]
}

func findMaximumXORSlow(nums []int) int {
	var max int
	var maxIxs []int
	for ix, n := range nums {
		if c := countBits(n); c > max {
			maxIxs = []int{ix}
			max = c
		} else if c == max {
			maxIxs = append(maxIxs, ix)
		}
	}

	var maXor int
	for _, ix := range maxIxs {
		m := nums[ix]
		for j, n := range nums {
			if ix == j || m == n {
				continue
			}

			if xor := m ^ n; xor > maXor {
				maXor = xor
			}
		}
	}

	return maXor
}

func countBits(n int) (c int) {
	for n > 0 {
		n >>= 1
		c++
	}
	return
}
