package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	t := new(trie)
	var instruction, str string
	for ; N > 0; N-- {
		fmt.Scan(&instruction, &str)
		switch instruction {
		case "find":
			fmt.Println(t.find(str))
		case "add":
			t.add(str)
		default:
			panic("instruction not found " + instruction)
		}
	}
}

type trie struct {
	count int
	tree  [26]*trie
}

func (t *trie) add(str string) {
	l := len(str)
	if l == 0 {
		return
	}

	for ; l > 0; l = len(str) {
		pos := str[0] - 'a'
		if t.tree[pos] == nil {
			t.tree[pos] = new(trie)
		}
		t.tree[pos].count++

		str = str[1:]
		t = t.tree[pos]
	}
}

func (t *trie) find(str string) int {
	l := len(str)
	if l == 0 {
		return 0
	}

	for ; l > 0; l = len(str) {
		pos := str[0] - 'a'
		if t.tree[pos] == nil {
			return 0
		}

		if l == 1 {
			return t.tree[pos].count
		}
		str = str[1:]
		t = t.tree[pos]
	}

	return 0
}
