package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abcd", 2))
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
	fmt.Println(removeDuplicates("pbbcggttciiippooaais", 2))
}

type pair struct {
	ch    byte
	count int
}

func (p *pair) String() string {
	return fmt.Sprintf("[%s %d]", string(p.ch), p.count)
}

func removeDuplicates(s string, k int) string {

	var stack []*pair
	for ix := range s {
		if len(stack) == 0 {
			stack = append(stack, &pair{ch: s[ix], count: 1})
			continue
		}
		top := stack[len(stack)-1]
		if top.ch == s[ix] {
			if top.count == k-1 {
				stack = stack[:len(stack)-1]
			} else {
				top.count++
			}
		} else {
			stack = append(stack, &pair{ch: s[ix], count: 1})
		}
	}
	// fmt.Println(stack)

	var res []byte
	for ix := range stack {
		for c := 0; c < stack[ix].count; c++ {
			res = append(res, stack[ix].ch)
		}
	}
	return string(res)
}

func removeDuplicatesSlow(s string, k int) string {
	ls := []byte(s)

	var changed bool
	for ix := 0; ix < len(ls)-1; ix++ {
		if ls[ix] == ls[ix+1] {
			start, end := ix, ix
			for end < len(ls)-1 && ls[end] == ls[end+1] && end < start+k {
				end++
			}
			if end+1 == start+k {
				ls = append(ls[:start], ls[end+1:]...)
				changed = true
			}
		}
	}
	if changed {
		return removeDuplicates(string(ls), k)
	}
	return string(ls)
}
