package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates("abbaca"))    // ca
	fmt.Println(removeDuplicates("azxxzy"))    // ay
	fmt.Println(removeDuplicates("aaaaaaaaa")) // a
}

func removeDuplicates(s string) string {
	var stack []byte

	for ix := range s {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			if top == s[ix] {
				stack = stack[:len(stack)-1] // pop
				// fmt.Println(stack)
			} else {
				stack = append(stack, s[ix])
			}
		} else {
			stack = append(stack, s[ix])
		}
	}

	// var b strings.Builder
	// for ix := 0; ix < len(stack); ix++ {
	// 	b.WriteByte(stack[ix])
	// }
	// return b.String()
	return string(stack)
}
