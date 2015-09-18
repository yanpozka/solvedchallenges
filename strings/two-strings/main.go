package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	read := bufio.NewReader(os.Stdin)

	line, _ := read.ReadString('\n')
	T, _ := strconv.Atoi(line[0 : len(line)-1])

	for ; T > 0; T-- {
		letters := make(map[rune]bool)

		line, _ = read.ReadString('\n')
		A := line[0 : len(line)-1]

		line, _ = read.ReadString('\n')
		B := line[0 : len(line)-1]

		for _, r := range A {
			letters[r] = true
		}

		var ok bool
		for _, r := range B {
			if _, contains := letters[r]; contains {
				ok = true
				break
			}
		}
		if ok {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
