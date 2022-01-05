package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var n uint32 = 43261596
	// fmt.Printf("%032b\n", n)
	// n >>= 1
	// fmt.Printf("%032b\n", n)

	fmt.Printf("%032b\n%032b\n", 43261596, reverseBits(43261596))
	fmt.Printf("%032b\n%032b\n", 4294967293, reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	bn := fmt.Sprintf("%032b", num)
	data := make([]byte, 32)

	var c int
	for ix := len(bn) - 1; ix >= 0; ix-- {
		data[c] = bn[ix]
		c++
	}

	res, _ := strconv.ParseUint(string(data), 2, 32)
	return uint32(res)
}
