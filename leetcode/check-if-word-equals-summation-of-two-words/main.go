package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(toInt64("acb"))
	fmt.Println(toInt64("cba"))
	fmt.Println(toInt64("cdb"))

	fmt.Println(isSumEqual("acb", "cba", "cdb"))
	fmt.Println(isSumEqual("aaa", "a", "aab"))
	fmt.Println(isSumEqual("aaa", "a", "aaaa"))
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return toInt64(firstWord)+toInt64(secondWord) == toInt64(targetWord)
}

func toInt64(str string) int64 {
	var b strings.Builder
	for ix := range str {
		val := strconv.Itoa(int(str[ix] - 'a'))
		b.WriteString(val)
	}
	n, _ := strconv.ParseInt(b.String(), 10, 64)
	return n
}

func toInt64Base10(str string) int64 {
	var num int64
	exp := len(str) - 1
	for ix := range str {
		val := int64(str[ix] - 'a')
		num += val * int64(math.Pow10(exp))
		exp--
	}
	return num
}
