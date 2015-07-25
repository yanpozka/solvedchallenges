package main

import (
	"fmt"
	"math/big"
)

func main() {
	var A, B, N int64

	fmt.Scanf("%d%d%d", &A, &B, &N)

	f, s := big.NewInt(A), big.NewInt(B)

	for ; N > 2; N-- {
		tmp := big.NewInt(int64(0))
		tmp.Abs(s)
		s = s.Mul(s, s)
		s = s.Add(s, f)
		f = tmp
	}
	fmt.Println(s)
}
