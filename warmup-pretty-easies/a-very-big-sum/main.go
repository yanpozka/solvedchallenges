package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	var sum, elewa int64
	for ; N > 0; N-- {
		fmt.Scanf("%d", &elewa)
		sum += elewa
	}
	fmt.Println(sum)
}
