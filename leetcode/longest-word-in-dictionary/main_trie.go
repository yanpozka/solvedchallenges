package main

import "fmt"

func main() {
	in := []string{"w", "wo", "wor", "worl", "world"}
	fmt.Println("longest word:", longestWord(in))

	in = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println("longest word:", longestWord(in))

	in = []string{"ts", "x", "e", "pbhj", "opto", "xhigy", "erikz", "pbh", "opt", "erikzb", "eri", "erik", "xlye", "xhig", "optoj", "optoje", "xly", "pb", "xhi", "x", "o"}
	fmt.Println("longest word:", longestWord(in))
}

type trie struct {
	chs [26]*trie
	end bool
}

func longestWord(words []string) string {
	root := new(trie)
	for _, w := range words {
		insertWord(root, w)
	}

	type item struct {
		node *trie
		prev string
	}
	q := []*item{{node: root}}
	var max string
	for len(q) > 0 {
		curr := q[0] // top
		q = q[1:]    // pop

		for ix, ch := range curr.node.chs {
			if ch == nil || !ch.end {
				continue
			}

			word := curr.prev + string('a'+ix)
			if len(word) > len(max) {
				max = word
			} else if len(word) == len(max) {
				if word < max {
					max = word
				}
			}
			q = append(q, &item{
				node: ch, prev: word,
			})
		}
	}
	return max
}

func insertWord(t *trie, w string) {
	curr := t
	for ix := 0; ix < len(w); ix++ {
		pos := w[ix] - 'a'
		if curr.chs[pos] == nil {
			curr.chs[pos] = new(trie)
		}
		curr = curr.chs[pos]
	}
	curr.end = true
}
