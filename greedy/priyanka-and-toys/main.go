package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	ws := make([]int, N)

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		c, _ := strconv.Atoi(scanner.Text())
		ws[ix] = c
	}

	sort.Ints(ws)

	var count, inc int
	for ix := 0; ix < N; ix += inc {
		inc = 0
		for k := ix; k < N && ws[k] <= ws[ix]+4; k++ {
			inc++
		}
		count++
	}
	fmt.Println(count)
}
