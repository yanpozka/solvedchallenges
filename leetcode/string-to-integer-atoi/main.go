package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("    +012"))
	fmt.Println(myAtoi("21474836460"))
	fmt.Println(myAtoi("20000000000000000000"))
	fmt.Println(myAtoi("  0000000000012345678"))
}

func myAtoi(str string) int {
	var s string
	//fmt.Println(strings.TrimSpace(str))
	for ix, r := range strings.TrimSpace(str) {
		if r >= '0' && r <= '9' {
			s += string(r)
		} else if ix == 0 {
			s += string(r)
		} else {
			break
		}
	}
	//fmt.Printf("%q\n", s)

	n64, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		if n64 < math.MinInt32 {
			return int(math.MinInt32)
		}
		if n64 > math.MaxInt32 {
			return int(math.MaxInt32)
		}
		return int(n64)
	}

	i := new(big.Int)
	if _, err := fmt.Sscan(s, i); err != nil {
		//fmt.Println(err)
		return 0
	}
	minI := big.NewInt(int64(math.MinInt32))
	maxI := big.NewInt(int64(math.MaxInt32))

	if d := i.Cmp(minI); d == -1 {
		return int(math.MinInt32)
	}
	if d := i.Cmp(maxI); d == 1 {
		return int(math.MaxInt32)
	}

	return int(i.Int64())
}
