package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var index int
	var first_char byte

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		current_word := scanner.Bytes()
		first_char = current_word[0]
		index = -1

		for ix := len(current_word) - 1; ix > 1; ix-- {
			if current_word[ix] > current_word[ix-1] {
				index = ix
				break
			}
		}

		if index == -1 {
			if first_char < current_word[1] {
				var min byte = 255
				var ix_min int = -1

				for i, c := range current_word {
					if c > first_char && c-first_char < min {
						min = c - first_char
						ix_min = i
					}
				}
				var mx byte = current_word[ix_min]
				p1 := current_word[:ix_min]

				if ix_min+1 < len(current_word) {
					p1 = append(p1, current_word[ix_min+1:]...)
				}
				var rasta ByteSlice = p1
				sort.Sort(rasta)

				fmt.Printf("%s%s\n", string(mx), string(rasta))

			} else {
				fmt.Println("no answer")
			}
		} else {
			current_word[index], current_word[index-1] = current_word[index-1], current_word[index]
			var rabo ByteSlice = current_word[index:]
			sort.Sort(rabo)

			fmt.Printf("%s%s\n", string(current_word[:index]), string(rabo))
		}
	}
}
