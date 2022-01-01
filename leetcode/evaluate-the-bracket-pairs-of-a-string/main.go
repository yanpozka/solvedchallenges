package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(evaluate("(name)is(age)yearsold", [][]string{
		{"name", "bob"}, {"age", "two"},
	}))
	fmt.Println(evaluate("hi(name) and hi(name)", [][]string{{"a", "b"}}))
	fmt.Println(evaluate("(a)(a)(a)aaa(b)(a)", [][]string{{"a", "yes"}}))
}

func evaluate(s string, knowledge [][]string) string {
	data := map[string]string{}
	for _, p := range knowledge {
		data[p[0]] = p[1]
	}

	b := new(strings.Builder)
	var prevIx int
	for ix := 0; ix < len(s); ix++ {
		if s[ix] != '(' {
			continue
		}

		b.WriteString(s[prevIx:ix])
		ix++
		start := ix
		for ; ix < len(s); ix++ {
			if s[ix] == ')' {
				break
			}
		}
		prevIx = ix + 1
		key := s[start:ix]
		if val, found := data[key]; found {
			b.WriteString(val)
		} else {
			b.WriteString("?")
		}
	}
	b.WriteString(s[prevIx:])
	return b.String()
}

func evaluateSlow(s string, knowledge [][]string) string {
	data := map[string]string{}
	for _, p := range knowledge {
		data[p[0]] = p[1]
	}
	checked := map[string]bool{}

	sol := s
	re := regexp.MustCompile(`\(([a-z]+)\)`)
	for _, key := range re.FindAllString(s, -1) {
		trimKey := key[1 : len(key)-1]
		if checked[trimKey] {
			continue
		}
		checked[trimKey] = true
		if val, found := data[trimKey]; !found {
			sol = strings.ReplaceAll(sol, key, "?")
		} else {
			sol = strings.ReplaceAll(sol, key, val)
		}
	}

	return sol
}
