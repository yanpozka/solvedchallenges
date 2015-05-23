package main

import "fmt"

func main() {
	var line string
	letters_count := make(map[byte]int)

	fmt.Scanf("%s", &line)
	for ix := range line {
		letters_count[line[ix]]++
	}
	var odd_count int = 0
	for l := range letters_count {
		if letters_count[l]%2 != 0 {
			odd_count++
		}
	}

	if odd_count > 1 {
		fmt.Println("NO")
	} else {
		if len(line)%2 != 0 {
			if odd_count == 1 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		} else {
			if odd_count == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}

}
