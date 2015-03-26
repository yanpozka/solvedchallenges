package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var counts [100]uint64

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var parts []string = strings.Split(scanner.Text(), " ")

		for ix := range parts {
			if parts[ix] == "" {
				continue
			}
			num, _ := strconv.Atoi(parts[ix])
			counts[num]++
		}
	}

	for ix := range counts {
		fmt.Print(counts[ix])
		if ix < 99 {
			fmt.Print(" ")
		}
	}
}
