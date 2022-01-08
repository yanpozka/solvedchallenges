package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findDifferentBinaryString([]string{"0"}))
	fmt.Println(findDifferentBinaryString([]string{"1"}))
	fmt.Println(findDifferentBinaryString([]string{"10", "11"}))
	fmt.Println(findDifferentBinaryString([]string{"01", "10"}))
	fmt.Println(findDifferentBinaryString([]string{"00", "01"}))
	fmt.Println(findDifferentBinaryString([]string{"111", "011", "001"}))
}

func findDifferentBinaryString(nums []string) string {
	bitSize := len(nums)
	if bitSize == 1 {
		if nums[0] == "1" {
			return "0"
		}
		return "1"
	}
	set := map[int]bool{}

	for _, nStr := range nums {
		num, _ := strconv.ParseInt(nStr, 2, 16)
		set[int(num)] = true
	}

	if _, found := set[0]; !found {
		var res string
		for bitSize > 0 {
			res += "0"
			bitSize--
		}
		return res
	}

	start := 1 << (bitSize - 1)
	end := 1 << bitSize

	for n := start; n < end; n++ {
		if _, found := set[n]; !found {
			return strconv.FormatInt(int64(n), 2)
		}
	}
	return ""
}
