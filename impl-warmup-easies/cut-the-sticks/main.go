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
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	nums := make([]int, len(parts))

	var min int = 100000
	for ix, s := range parts {
		nums[ix], _ = strconv.Atoi(s)
		if nums[ix] < min {
			min = nums[ix]
		}
	}

	for {
		var count, new_min int = 0, 1000000

		for ix, n := range nums {
			if n > 0 {
				nums[ix] -= min

				if nums[ix] > 0 && nums[ix] < new_min {
					new_min = nums[ix]
				}
				count++
			}
		}

		if count > 0 {
			fmt.Println(count)
			min = new_min
		} else {
			break
		}
	}
}
