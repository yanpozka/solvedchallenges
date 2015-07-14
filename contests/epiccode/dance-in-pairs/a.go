package main

import "fmt"

func main() {
	fmt.Println(100000, 1000000000)

	for i := 0; i < 100000; i++ {
		fmt.Printf("%d", i)
		if i < 100000-1 {
			fmt.Print(" ")
		}
	}

	for i := 0; i < 100000; i++ {
		fmt.Printf("%d", i)
		if i < 100000-1 {
			fmt.Print(" ")
		}
	}

}
