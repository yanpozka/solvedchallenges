package main

import "fmt"

func main() {
	var line string
	fmt.Scan(&line)

	var t, result int
	for _, c := range line {
		switch t {
		case 0, 2:
			if c != 'S' {
				result++
			}
		case 1:
			if c != 'O' {
				result++
			}
		}
		t++
		t %= 3
	}

	fmt.Println(result)
}
