package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	an := new(big.Int)
	fmt.Sscanf(a, "%b", an)
	bn := new(big.Int)
	fmt.Sscanf(b, "%b", bn)
	return fmt.Sprintf("%b", an.Add(an, bn))
}
