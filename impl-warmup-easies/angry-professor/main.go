package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0; T-- {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		K, _ := strconv.Atoi(scanner.Text())

		count := 0
		for ix := 0; ix < N && scanner.Scan(); ix++ {
			n, _ := strconv.Atoi(scanner.Text())
			if n <= 0 {
				count++
			}
		}
		if count < K {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
