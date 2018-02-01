// TODO: it's missing inner parentheses !!

//
// https://leetcode.com/problems/generate-parentheses/description/
//
package main

import "fmt"

func main() {
	// ss := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}

	ss := []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}

	res := generateParenthesis(5)
	m := map[string]bool{}
	for _, s := range res {
		m[s] = true
	}
	for _, str := range ss {
		if _, ok := m[str]; !ok {
			fmt.Println(str)
		}
	}
	fmt.Println(ss)
	fmt.Println(res)
}

type pair struct {
	str    string
	isSame bool
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}

	strs := generatePP(n)
	res := make([]string, len(strs))
	for ix, s := range strs {
		res[ix] = s.str
	}
	return res
}

func generatePP(n int) []*pair {
	if n == 2 {
		return []*pair{
			{"(())", false},
			{"()()", true},
		}
	}

	res := make([]*pair, 0, n)
	others := generatePP(n - 1)

	for _, o := range others {
		res = append(res, &pair{str: "(" + o.str + ")"})
		res = append(res, &pair{str: "()" + o.str, isSame: true})
		if !o.isSame {
			res = append(res, &pair{str: o.str + "()"})
		}
	}
	return res
}
