package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Al = "a"[0]

type tries struct {
	count int
	tree  [10]*tries
}

func add(t *tries, cad string) bool {
	if cad == "" {
		return true
	}

	var pos int = int(cad[0] - Al)

	if t.tree[pos] == nil {
		t.tree[pos] = &tries{}
	} else {
		if len(cad) == 1 || t.tree[pos].count >= 1 {
			return false
		}
	}

	if len(cad) == 1 {
		t.tree[pos].count++
	}
	return add(t.tree[pos], cad[1:])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	t := &tries{}

	for ix := 0; ix < N && scanner.Scan(); ix++ {
		s := scanner.Text()
		if add(t, s) == false {
			fmt.Println("BAD SET")
			fmt.Println(s)
			return
		}
	}

	fmt.Println("GOOD SET")
}
