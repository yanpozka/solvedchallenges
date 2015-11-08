package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		R, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		// C, _ := strconv.Atoi(scanner.Text())

		lines := make([]string, R)

		for ix := 0; ix < R && scanner.Scan(); ix++ {
			lines[ix] = scanner.Text()
		}

		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		// n, _ := strconv.Atoi(scanner.Text())

		tosearch := make([]string, m)

		for ix := 0; ix < m && scanner.Scan(); ix++ {
			tosearch[ix] = scanner.Text()
		}

		var found bool

		for ix := 0; ix < R-m; ix++ {

			if found_index := strings.Index(lines[ix], tosearch[0]); found_index != -1 {
				b := 1
				for a := ix + 1; b < m; b++ {
					if currix := strings.Index(lines[a], tosearch[b]); currix != found_index {
						if strings.Index(lines[a][found_index:], tosearch[b]) != 0 {
							break
						}
					}
					a++
				}

				if b == m {
					found = true
					break
				}
			}
		} // for

		if found {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
