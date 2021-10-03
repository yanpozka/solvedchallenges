package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	freq := map[int32]int32{}
	counts := map[int32]int{}
	var sol []int32

	for _, q := range queries {
		switch q[0] {
		case 1:
			if orgC := freq[q[1]]; orgC > 0 {
				counts[orgC]--
				if c := counts[orgC]; c == 0 {
					delete(counts, orgC)
				}
			}
			freq[q[1]]++
			counts[freq[q[1]]]++
		case 2:
			if _, ok := freq[q[1]]; ok {
				if orgC := freq[q[1]]; orgC > 0 {
					counts[orgC]--
					if c := counts[orgC]; c == 0 {
						delete(counts, orgC)
					}
				}
				freq[q[1]]--
				if c := freq[q[1]]; c == 0 {
					delete(freq, q[1])
				} else {
					counts[c]++
				}
			}
		case 3:
			z := q[1]
			if _, found := counts[z]; found {
				sol = append(sol, 1)
			} else {
				sol = append(sol, 0)
			}
		}
	}

	return sol
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

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
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
