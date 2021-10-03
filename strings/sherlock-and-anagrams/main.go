package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sherlockAndAnagrams(s string) int32 {
	var count int32

	for l := 1; l < len(s); l++ {
		var subs []string
		for ix := 0; ix < len(s); ix++ {
			lim := ix + l
			if lim >= len(s) {
				lim = len(s)
			}
			subs = append(subs, s[ix:lim])
		}

		for ix := range subs {
			for k := ix + 1; k < len(subs); k++ {
				if areAnagrams(subs[ix], subs[k]) {
					count++
				}
			}
		}
	}

	return count
}

func areAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return true
	}
	var aC [26]int
	var bC [26]int

	for ix := range a {
		aC[a[ix]-'a']++
		bC[b[ix]-'a']++
	}

	for ix := range aC {
		if aC[ix] != bC[ix] {
			return false
		}
	}

	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
