package main

import "fmt"

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
}

func diStringMatch(str string) []int {
	s, e := 0, len(str)
	sol := make([]int, 0, len(str)+1)

	for _, ch := range str {
		if ch == 'I' {
			sol = append(sol, s)
			s++
		} else {
			sol = append(sol, e)
			e--
		}
	}
	sol = append(sol, s)
	return sol
}
