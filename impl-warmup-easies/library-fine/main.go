package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var dR, mR, yR, day, month, year int
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")

	dR, _ = strconv.Atoi(parts[0])
	mR, _ = strconv.Atoi(parts[1])
	yR, _ = strconv.Atoi(parts[2])

	returned := time.Date(yR, time.Month(mR), dR, 0, 0, 0, 0, time.UTC)

	scanner.Scan()
	parts = strings.Split(scanner.Text(), " ")

	day, _ = strconv.Atoi(parts[0])
	month, _ = strconv.Atoi(parts[1])
	year, _ = strconv.Atoi(parts[2])

	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	if returned.After(expected) {
		if returned.Year() > expected.Year() {
			fmt.Println(10000)
		} else if mR > month {
			fmt.Println((mR - month) * 500)
		} else {
			fmt.Println((dR - day) * 15)
		}
	} else {
		fmt.Println(0)
	}
}
