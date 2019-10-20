package main

import (
	"fmt"
)

func main() {
	/*const base float64 = 26
	ns = make([]int, 0, 7)
	for exp := 0; exp < cap(ns); exp++ {
		ns = append(ns, int(math.Pow(base, float64(exp))))
	}
	fmt.Println(ns)*/

	fmt.Println(titleToNumber(""))
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("C"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	ns := [...]int{1, 26, 676, 17576, 456976, 11881376, 308915776}
	if s == "" {
		return 0
	}
	var res int
	const base float64 = 26
	var exp int
	for ix := len(s) - 1; ix >= 0; ix-- {
		res += int(s[ix]-'A'+1) * ns[exp]
		exp++
	}
	return res
}
