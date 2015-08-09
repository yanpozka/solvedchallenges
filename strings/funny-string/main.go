package main

import (
	"bufio"
	"fmt"
	"math"
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
		S := scanner.Text()
		leng := len(S)
		var ok bool = true

		for ix, r := 1, leng-2; ix < leng; ix++ {
			if math.Abs(float64(S[ix])-float64(S[ix-1])) != math.Abs(float64(S[r])-float64(S[r+1])) {
				ok = false
				break
			}
			r--
		}
		if ok {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}
}
