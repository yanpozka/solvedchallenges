//
// https://www.careercup.com/question?id=5181387768332288
//
// https://play.golang.org/p/GPhsh0O6uk
//
package main

import (
	"fmt"
)

/***

{
	"+5412312" -> [2, 3]
	"+1234567" -> [1]

	"john@gmail.com"    -> [0]
	"john@fb.com"       -> [0, 3]
	"dan@gmail.com"     -> [1]
	"john123@skype.com" -> [2]
}

***/

// Problem: find duplicated nodes
//
func main() {
	input := [][]string{
		{"John", "john@gmail.com", "john@fb.com"},
		{"Dan", "dan@gmail.com", "+1234567"},
		{"john123", "+5412312", "john123@skype.com"},
		{"john1985", "+5412312", "john@fb.com"},
	}

	ids := make(map[string]int)
	set := newDisjointSet(len(input))

	for ix, data := range input { // O(n)
		for _, d := range data[1:] { // ignore name;  O(k)
			if other, contains := ids[d]; contains {
				set.Merge(ix, other)
				fmt.Println("contains so merged:", other, "and", ix)
			} else {
				fmt.Println("first key:", d, "node:", ix)
				ids[d] = ix
			}
		}
	}

	fmt.Println(set.Sets())
}

type DisjointSet struct {
	size       int
	tree, rank []int
}

func newDisjointSet(n int) *DisjointSet {
	tree := make([]int, n)
	for ix := range tree {
		tree[ix] = ix
	}

	return &DisjointSet{
		size: n,
		tree: tree,
		rank: make([]int, n),
	}
}

func (d *DisjointSet) Sets() (result [][]int) {
	nodes := make(map[int][]int)

	for ix, p := range d.tree {
		if ix == p {
			nodes[p] = append(nodes[p], p)
		} else {
			pix := d.Search(ix)
			nodes[pix] = append(nodes[pix], ix)
		}
	}

	for _, v := range nodes {
		// fmt.Println(k, v)
		result = append(result, v)
	}

	return
}

func (d *DisjointSet) Search(o int) int {
	if d.tree[o] == o {
		return o
	}
	return d.Search(d.tree[o])
}

func (d *DisjointSet) Merge(a, b int) {
	if a == b {
		return
	}
	pA, pB := d.Search(a), d.Search(b)
	if pA == pB {
		return
	}

	if d.rank[pA] > d.rank[pB] {
		d.tree[pB] = pA
	} else {
		if d.rank[pA] == d.rank[pB] {
			d.rank[pB]++
		}
		d.tree[pA] = pB
	}
}
