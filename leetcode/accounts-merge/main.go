package main

import (
	"fmt"
	"sort"
)

func main() {
	in := [][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}
	// Output: [["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
	fmt.Println(`[["John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]`)
	fmt.Println(accountsMerge(in))
}

func accountsMerge(accounts [][]string) [][]string {
	initTree(len(accounts))
	m := make(map[string]int)

	for ix, person := range accounts {
		for k := 1; k < len(person); k++ {
			email := person[k]
			if p, contains := m[email]; contains {
				union(p, ix)
			}
			m[email] = ix
		}
	}

	ids := make(map[int]map[string]bool, len(accounts))

	for ix := range tree {
		p := find(ix)
		if _, contains := ids[p]; !contains {
			ids[p] = make(map[string]bool)
		}
		emails := ids[p]
		if ix == p {
			person := accounts[p]
			for k := 1; k < len(person); k++ {
				emails[person[k]] = true
			}
		} else {
			person := accounts[ix]
			for k := 1; k < len(person); k++ {
				emails[person[k]] = true
			}
		}
		ids[p] = emails
	}

	res := make([][]string, 0, len(ids))

	for ix, emails := range ids {
		es := make([]string, 0, len(emails))
		for email := range emails {
			es = append(es, email)
		}
		sort.Strings(es)
		res = append(res, append([]string{accounts[ix][0]}, es...))
	}

	return res
}

var tree, rank []int

func initTree(n int) {
	tree = make([]int, n)
	for ix := range tree {
		tree[ix] = ix
	}
	rank = make([]int, n)
}

func find(r int) int {
	tmp := tree[r]
	for {
		if tmp == tree[tmp] {
			break
		}
		tmp = tree[tmp]
	}
	return tmp
}

func union(x, y int) {
	pX, pY := find(x), find(y)
	if pX == pY {
		return
	}
	if rank[pX] > rank[pY] {
		tree[pY] = pX
	} else {
		if rank[pX] == rank[pY] {
			rank[pY]++
		}
		tree[pX] = pY
	}
}
