package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

var val = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var num int

	subs := func(ix *int, ch, a, b byte) {
		if *ix+1 < len(s) {
			if s[*ix+1] == a {
				num += val[a] - val[ch]
				*ix++
			} else if s[*ix+1] == b {
				num += val[b] - val[ch]
				*ix++
			} else {
				num += val[ch]
			}
		} else {
			num += val[ch]
		}
	}

	for ix := 0; ix < len(s); ix++ {
		ch := s[ix]
		switch ch {
		case 'I':
			subs(&ix, ch, 'V', 'X')
		case 'X':
			subs(&ix, ch, 'L', 'C')
		case 'C':
			subs(&ix, ch, 'D', 'M')
		default:
			num += val[ch]
		}
	}
	return num
}
