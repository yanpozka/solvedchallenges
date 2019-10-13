package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))

	fmt.Println(
		replaceWords([]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"))
}

func replaceWords(dict []string, sentence string) string {
	m := map[int]map[string]bool{}
	var maxDictLen int

	for _, word := range dict {
		l := len(word)
		if m[l] == nil {
			m[l] = make(map[string]bool, 1)
		}
		m[l][word] = true
		if l > maxDictLen {
			maxDictLen = l
		}
	}
	fmt.Println(maxDictLen)

	result := strings.Split(sentence, " ")

	for ix, word := range result {

		for l := 1; l <= maxDictLen && l < len(word); l++ {
			if _, found := m[l][word[:l]]; found {
				result[ix] = word[:l]
				break
			}
		}
	}

	return strings.Join(result, " ")
}
