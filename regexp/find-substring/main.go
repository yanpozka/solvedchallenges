package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regExpWord = regexp.MustCompile(`[[:word:]]+`)
var N, T int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	lines := list.New()

	for ; N > 0 && scanner.Scan(); N-- {
		var words []string = regExpWord.FindAllString(scanner.Text(), -1)
		lines.PushBack(words)
	}

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	for ; T > 0 && scanner.Scan(); T-- {
		var current_word string = scanner.Text()
		var count, curr_word_len int = 0, len(current_word)

		for elem := lines.Front(); elem != nil; elem = elem.Next() {
			sentence := elem.Value.([]string)

			for ix := range sentence {
				var word_sentence = sentence[ix]
				if len(word_sentence) > curr_word_len {
					ftix := strings.Index(word_sentence, current_word)

					if ftix > 0 && ftix < len(word_sentence)-curr_word_len {
						count++
					}
				}
			}
		}
		fmt.Println(count)
	} // for T
}
