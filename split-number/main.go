package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var regExp = regexp.MustCompile(`([0-9]+)[\-\ ]([0-9]+)[\-\ ]([0-9]+)`)
var N int

func main() {
	fmt.Scanf("%d", &N)
	reader := bufio.NewScanner(os.Stdin)

	for ; N > 0; N-- {
		reader.Scan()

		group := regExp.FindStringSubmatch(reader.Text())

		fmt.Printf("CountryCode=%s,LocalAreaCode=%s,Number=%s\n", group[1], group[2], group[3])
	}
}
