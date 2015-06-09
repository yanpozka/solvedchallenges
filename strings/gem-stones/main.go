package main

import "fmt"

const ACODE = 97

func main() {
	var N, i byte
	var line string
	fmt.Scanf("%d", &N)

	var Letters [26]byte

	for i = 0; i < N; i++ {
		fmt.Scanf("%s", &line)
		var visited [26]bool

		for k := range line {
			if !visited[line[k]-ACODE] {
				Letters[line[k]-ACODE] += 1
				visited[line[k]-ACODE] = true
			}
		}
	}
	var count int = 0
	for i = 0; i < 26; i++ {
		if Letters[i] == N {
			count++
		}
	}

	fmt.Println(count)
}
