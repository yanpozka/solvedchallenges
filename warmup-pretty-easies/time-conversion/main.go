package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var line string = scanner.Text()

	parts := strings.Split(line, ":")

	hour, _ := strconv.Atoi(parts[0])
	min, _ := strconv.Atoi(parts[1])

	sec, _ := strconv.Atoi(parts[2][:len(parts[2])-2])
	period := strings.ToUpper(parts[2][len(parts[2])-2:])

	if period == "PM" {
		if hour != 12 {
			hour += 12
			if hour == 24 {
				hour = 0
			}
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}
	fmt.Printf("%s:%s:%s\n", format(hour), format(min), format(sec))
}

func format(n int) string {
	var result string = strconv.Itoa(n)
	if n < 10 {
		result = "0" + result
	}
	return result
}
