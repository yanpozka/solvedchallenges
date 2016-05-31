package main

import "fmt"

func main() {
	var N, Q int
	var line string

	fmt.Scan(&N)
	t := new(tries)

	for ; N > 0; N-- {
		fmt.Scan(&line)
		t.add(line)
	}

	fmt.Scan(&Q)
	for ; Q > 0; Q-- {
		fmt.Scan(&line)
		fmt.Println(t.search(line))
	}
}

var Al = "a"[0]

type tries struct {
	count int
	tree  [26]*tries
}

func (t *tries) add(str string) {
	if str == "" || t == nil {
		return
	}
	pos := int(str[0] - Al)
	if t.tree[pos] == nil {
		t.tree[pos] = &tries{}
	}
	if len(str) == 1 {
		t.tree[pos].count++
	}
	t.tree[pos].add(str[1:])
}

func (t *tries) search(str string) int {
	if str == "" || t == nil {
		return 0
	}
	pos := int(str[0] - Al)
	if len(str) == 1 {
		return t.tree[pos].count
	}
	return t.tree[pos].search(str[1:])
}
