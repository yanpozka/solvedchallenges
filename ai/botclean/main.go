package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	DIRT = 100 // d
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	bot_row, _ := strconv.Atoi(parts[0])
	bot_col, _ := strconv.Atoi(parts[1])
	var min_dist float64 = 99999
	var min_row, min_col = -1, -1

	for i := 0; i < 5 && scanner.Scan(); i++ {
		line := scanner.Text()
		for k := range line {
			if line[k] == DIRT {
				if bot_row == i && bot_col == k {
					fmt.Println("CLEAN")
					os.Exit(0)
				}
				var cd float64 = dist(bot_row, bot_col, i, k)
				if cd < min_dist {
					min_dist = cd
					min_row, min_col = i, k
				}
			}
		}
	}
	if min_dist == 99999 {
		panic("No dirt found :(")
	}

	var first_direct string

	if min_col > bot_col {
		first_direct = "RIGHT"
	} else {
		if min_col < bot_col {
			first_direct = "LEFT"
		} else { // same col
			if min_row > bot_row {
				first_direct = "DOWN"
			} else {
				first_direct = "UP"
			}
		}
	}

	fmt.Println(first_direct)
}

func dist(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
}
