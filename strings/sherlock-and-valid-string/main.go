package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isValid(s string) string {
	if s == "" {
		return "NO"
	}
	var freqs [26]int

	for _, ch := range s {
		freqs[ch-'a']++
	}

	var numFail int
	prev := -1
	for ix := 0; ix < len(freqs); ix++ {
		if freqs[ix] == 0 {
			continue
		}
		// fmt.Println(freqs[ix], string('a'+ix))
		if prev != -1 {
			if prev != freqs[ix] {
				numFail++
				if numFail > 1 {
					return "NO"
				}
				// ixFail = ix
			}
		} else {
			prev = freqs[0]
		}
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
