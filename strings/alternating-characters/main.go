package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	var line string

	fmt.Scanf("%d", &N)

	reader := bufio.NewReader(os.Stdin)

	for ; N > 0; N-- {
		line, _ = reader.ReadString('\n')

		current_char := line[0]
		var current_index, count, size int = 0, 0, len(line)
		var flag_never_change bool = true

		for ix := 1; ix < size; ix++ {
			if current_char != line[ix] {
				count += (ix - current_index - 1)

				current_index = ix
				current_char = line[ix]

				flag_never_change = false
			} else {
				flag_never_change = true
			}
		}
		if flag_never_change {
			count += (size - current_index - 1)
		}

		fmt.Println(count)
	}
}
