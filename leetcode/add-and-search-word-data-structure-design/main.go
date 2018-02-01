//
// https://leetcode.com/problems/add-and-search-word-data-structure-design/description/
//
package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")

	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search(".at"))
	obj.AddWord("bat")
	fmt.Println(obj.Search(".at"))
	fmt.Println(obj.Search("an."))
	fmt.Println(obj.Search("a.d."))
	fmt.Println(obj.Search("b."))
	fmt.Println(obj.Search("a.d"))
	fmt.Println(obj.Search("."))
}

type WordDictionary struct {
	children [26]*WordDictionary
	isLeaf   bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	if word == "" {
		return
	}
	if this == nil {
		*this = Constructor()
	}
	t := this
	for ix := range word {
		l := word[ix] - 'a'

		if t.children[l] == nil {
			t.children[l] = new(WordDictionary)
		}
		if ix == len(word)-1 {
			t.children[l].isLeaf = true
			break
		}
		t = t.children[l]
	}
}

func (this *WordDictionary) Search(word string) bool {
	if this == nil || word == "" {
		return false
	}

	if word[0] == '.' {
		if len(word) == 1 {
			for _, ch := range this.children {
				if ch != nil && ch.isLeaf {
					return true
				}
			}
			return false
		}
		for _, ch := range this.children {
			if ch.Search(word[1:]) {
				return true
			}
		}
		return false
	}

	l := word[0] - 'a'
	if this.children[l] == nil {
		return false
	}
	if len(word) == 1 && this.children[l].isLeaf {
		return true
	}
	return this.children[l].Search(word[1:])
}
