// https://play.golang.org/p/OPANh6QnDY
//
package main

import (
	"fmt"
	"strings"
)

func main() {
	dict := []string{"cat", "bat", "rat"}
	const s = "the cattle was rattled by the battery"

	fmt.Println(replaceWords(dict, s))
}

func replaceWords(dict []string, sentence string) string {
	tree := newTrie()

	words := strings.Split(sentence, " ")

	for _, w := range dict {
		tree.add(w)
	}

	for ix, w := range words {
		ok, p := tree.hasPrefix(w)
		if ok {
			words[ix] = p
		}
	}

	return strings.Join(words, " ")
}

type trie struct {
	ch     []*trie
	isWord bool
}

func newTrie() *trie {
	return &trie{ch: make([]*trie, 26)}
}

func (this *trie) add(str string) {
	tmp := this
	for i := range str {
		index := str[i] - 'a'
		if tmp.ch[index] == nil {
			tmp.ch[index] = newTrie()
		}
		if i == len(str)-1 {
			tmp.ch[index].isWord = true
		}
		tmp = tmp.ch[index]
	}
}

func (this *trie) hasPrefix(in string) (bool, string) {
	tmp := this
	for i := range in {
		index := in[i] - 'a'
		if tmp.ch[index] == nil {
			return false, ""
		}
		if tmp.ch[index].isWord {
			return true, in[:i+1]
		}
		tmp = tmp.ch[index]
	}
	return false, ""
}
