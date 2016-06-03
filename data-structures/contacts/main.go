package main

import "fmt"

func main() {
	var N int
	var action, line string

	fmt.Scan(&N)
	t := new(tries)

	for ; N > 0; N-- {
		fmt.Scan(&action, &line)
		if action == "add" {
			t.add(line)
		} else {
			fmt.Println(t.search(line))
		}
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
		t.tree[pos] = &tries{count: 1}
	} else {
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
		if st := t.tree[pos]; st != nil {
			return st.count
		}
		return 0
	}
	return t.tree[pos].search(str[1:])
}
