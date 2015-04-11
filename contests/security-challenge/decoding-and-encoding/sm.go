package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

var N int

func main() {

	fmt.Scanf("%d", &N)

	for ; N > 0; N-- {
		var line string
		fmt.Scanf("%s", &line)
		origin, _ := base32.StdEncoding.DecodeString(line)

		r := base64.StdEncoding.EncodeToString(origin)
		fmt.Println(string(r))
	}
}
