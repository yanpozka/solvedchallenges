package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var regEx = regexp.MustCompile(`http://([a-z0-9]+[.]{1}[a-z0-9\-\.]*[a-z0-9]+)+`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteString('\n')
	}

	urls_map := make(map[string]bool)

	all_groups := regEx.FindAllStringSubmatch(buffer.String(), -1)

	for i := range all_groups {
		group := all_groups[i]

		fmt.Println(group)

		if len(group) > 1 {
			// fmt.Println(group)
			urls_map[group[1]] = true
		}
	}

	var size int = len(urls_map)
	var urls sort.StringSlice = make(sort.StringSlice, 0, size)
	for ur := range urls_map {
		if strings.Index(ur, "www") == 0 || strings.Index(ur, "ww2") == 0 {
			ur = ur[4:]
		}
		urls = append(urls, ur)
	}
	urls.Sort()
	size--

	for ix := range urls {
		fmt.Print(urls[ix])
		if ix < size {
			fmt.Print(";")
		}
	}
}
