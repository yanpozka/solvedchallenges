package main

import (
	"bufio"
	"fmt"
	"os"
)

const Acode = 97

func main() {
	reader := bufio.NewReader(os.Stdin)
	var line string
	var n int
	fmt.Scanf("%d", &n)

	line, _ = reader.ReadString('\n')
	line = line[:len(line)-1]
	counts := make([]int, 26)

	for _, c := range line {
		counts[c-Acode] += 1
	}

	var sum int
	for ix := range counts {
		if counts[ix] > 0 {
			sum += (counts[ix] * (counts[ix] + 1)) / 2
		}
	}
	fmt.Println(sum)
}
