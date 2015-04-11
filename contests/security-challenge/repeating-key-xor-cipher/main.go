package main

import (
	"fmt"
	"strconv"
)

func main() {
	var key, cyphred string

	fmt.Scanf("%s\n%s", &key, &cyphred)
	var len_key, size_cyphred, current int = len(key), len(cyphred), 0
	var buffer []byte = make([]byte, size_cyphred/2)

	for ix, k := 0, 0; ix < size_cyphred; ix += 2 {
		cr, _ := strconv.ParseInt(cyphred[ix:ix+2], 16, 8)
		buffer[current] = key[k] ^ uint8(cr)
		current++
		k++
		if k == len_key {
			k = 0
		}
	}
	fmt.Println(string(buffer))
}
