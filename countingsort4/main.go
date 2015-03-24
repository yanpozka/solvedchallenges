package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wd struct {
	Word     string
	IsShowed bool
}

func main() {
	nums_map := make(map[int][]Wd)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	N, _ := strconv.Atoi(scanner.Text())

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		parts := strings.Split(scanner.Text(), " ")
		c_n, _ := strconv.Atoi(parts[0])

		w := Wd{Word: parts[1], IsShowed: ix < N/2}
		if list, exist := nums_map[c_n]; exist {
			nums_map[c_n] = append(list, w)
		} else {
			l := append(make([]Wd, 0, 100), w)
			nums_map[c_n] = l

		}
	}

	for ix := 0; ix < N; ix++ {
		if list, exist := nums_map[ix]; exist {

			for k := range list {
				if !list[k].IsShowed {
					fmt.Printf("%s ", list[k].Word)
				} else {
					fmt.Print("- ")
				}
			}

		}
	}
}
