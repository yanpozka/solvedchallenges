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
	var index int
	var first_char byte

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		current_word := scanner.Bytes()
		index = -1

		for ix := len(current_word) - 1; ix > 0; ix-- {
			if current_word[ix] > current_word[ix-1] {
				index = ix
				break
			}
		}

		if index != -1 {
			lencw := len(current_word)

			if lencw > 1 {
				var min byte = 255
				var ix_min int = -1

				first_char = current_word[index-1]

				for i := index; i < lencw; i++ {
					c := current_word[i]
					if c > first_char && c-first_char < min {
						min = c - first_char
						ix_min = i
					}
				}

				current_word[index-1], current_word[ix_min] = current_word[ix_min], current_word[index-1]

				var sapo ByteSlice = current_word[index:]
				sort.Sort(sapo)

				fmt.Printf("%s%s\n", string(current_word[:index]), string(sapo))

			} else {
				fmt.Println("no answer")
			}
		} else {
			fmt.Println("no answer")
		}
	}
}

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
