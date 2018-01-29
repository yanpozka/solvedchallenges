//
// https://leetcode.com/problems/implement-magic-dictionary/description/
// https://play.golang.org/p/kPpqzxOx_oY
package main

import "fmt"

func main() {
	d := Constructor()
	d.BuildDict([]string{"hello", "leetcode", ""})

	fmt.Println(d.Search(""))
	fmt.Println(d.Search("hello"))
	fmt.Println(d.Search("hhllo"))
	fmt.Println(d.Search("hell"))
	fmt.Println(d.Search("leetcoded"))
	fmt.Println(d.Search("leetcode"))
}

type MagicDictionary struct {
	set map[string]bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{set: make(map[string]bool, 128)}
}

func (this *MagicDictionary) BuildDict(dict []string) {
	for _, word := range dict {
		for ix := range word {
			ch := word[ix] - 'a'
			for l := 'a'; l <= 'z'; l++ {
				if byte(l-'a') == ch {
					continue
				}
				this.set[fmt.Sprintf("%s%s%s", word[:ix], string(l), word[ix+1:])] = true
			}
		}
	}
}

func (this *MagicDictionary) Search(word string) bool {
	return this.set[word]
}
