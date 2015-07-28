package main

import (
	"fmt"
	"math/big"
)

func main() {
	var N int64
	fmt.Scanf("%d", &N)

	result := big.NewInt(N)

	for ix := N - 1; ix > 1; ix-- {
		result = result.Mul(result, big.NewInt(ix))
	}

	fmt.Println(result)
}
