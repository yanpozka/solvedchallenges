package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'contacts' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts 2D_STRING_ARRAY queries as parameter.
 */

func contacts(queries [][]string) []int32 {
	// Write your code here
	prefixes := map[string]int32{}
	var sol []int32

	for _, query := range queries {
		word := query[1]
		if query[0] == "add" {
			addPrefixes(prefixes, word)
		} else if query[0] == "find" {
			sol = append(sol, prefixes[word])
		}
	}

	return sol
}

func addPrefixes(m map[string]int32, word string) {
	if word == "" {
		return
	}
	for size := 1; size <= len(word); size++ {
		m[word[:size]]++
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	queriesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries [][]string
	for i := 0; i < int(queriesRows); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []string
		for _, queriesRowItem := range queriesRowTemp {
			queriesRow = append(queriesRow, queriesRowItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := contacts(queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
