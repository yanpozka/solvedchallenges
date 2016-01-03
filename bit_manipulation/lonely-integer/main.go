package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Scan()

	parts := strings.Split(scanner.Text(), " ")

	nums := make(map[string]int)

	for _, s := range parts {
		if _, ok := nums[s]; !ok {
			nums[s] = 0
		}
		nums[s] += 1
	}
	for _, s := range parts {
		if c := nums[s]; c == 1 {
			fmt.Println(s)
		}
	}
}
