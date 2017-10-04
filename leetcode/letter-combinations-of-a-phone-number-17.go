// https://play.golang.org/p/_ItUfr6NxA
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", letterCombinations("23"))
}

var dict = [][]string{
    {}, // 0
    {}, // 1
    {"a", "b", "c"}, // 2
    {"d", "e", "f"}, // 3
    {"g", "h", "i"}, // 4
    {"j", "k", "l"}, // 5
    {"m", "n", "o"}, // 6
    {"p", "q", "r", "s"},// 7
    {"t", "u", "v"},     // 8
    {"w", "x", "y", "z"},// 9
}

func letterCombinations(digits string) []string {
    if digits == "" {
        return nil
    }
    if len(digits) == 1 {
        return dict[digits[0] - '0']
    }
    
    var result []string
    for _, d := range dict[int(digits[0] - '0')] {
        for _, other := range letterCombinations(digits[1:]) {
            result = append(result, d + other)
        }
    }
    return result
}
