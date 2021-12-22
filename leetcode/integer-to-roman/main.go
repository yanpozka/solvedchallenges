package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994)) // MCMXCIV
	fmt.Println(intToRoman(3350)) // MMMCCCL
	fmt.Println(intToRoman(3999)) // MMMCMXCIX
}

func intToRoman(num int) string {
	numStr := strconv.Itoa(num)

	var res string
	var c int
	for ix := len(numStr) - 1; ix >= 0; ix-- {
		res = chars[numStr[ix]][c] + res
		c++
	}

	return res
}

var chars = map[byte]map[int]string{
	'1': {
		0: "I",
		1: "X",
		2: "C",
		3: "M",
	},
	'2': {
		0: "II",
		1: "XX",
		2: "CC",
		3: "MM",
	},
	'3': {
		0: "III",
		1: "XXX",
		2: "CCC",
		3: "MMM",
	},
	'4': {
		0: "IV",
		1: "XL",
		2: "CD",
	},
	'5': {
		0: "V",
		1: "L",
		2: "D",
	},
	'6': {
		0: "VI",
		1: "LX",
		2: "DC",
	},
	'7': {
		0: "VII",
		1: "LXX",
		2: "DCC",
	},
	'8': {
		0: "VIII",
		1: "LXXX",
		2: "DCCC",
	},
	'9': {
		0: "IX",
		1: "XC",
		2: "CM",
	},
}
