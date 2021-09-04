package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// all trues
	for _, s := range []string{"0", ".1", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"} {
		fmt.Println(s, isNumber(s))
	}
	println()

	// all falses
	for _, s := range []string{"6+1", ".", "e", "0..", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"} {
		fmt.Println(s, isNumber(s))
	}
}

func isNumber(s string) bool {
	parts := strings.Split(s, "e")
	if len(parts) > 2 {
		return false
	}
	if len(parts) == 2 {
		return checkExpNum(parts)
	}
	parts = strings.Split(s, "E")
	if len(parts) > 2 {
		return false
	}
	if len(parts) == 2 {
		return checkExpNum(parts)
	}

	var dashC, plusC, dotC int
	for _, c := range s {
		if c == '-' {
			dashC++
			if dashC > 2 {
				return false
			} else {
				continue
			}
		}
		if c == '+' {
			plusC++
			if plusC > 2 {
				return false
			} else {
				continue
			}
		}
		if c == '.' {
			dotC++
			if dotC > 1 {
				return false
			} else {
				continue
			}
		}
		if c < '0' || c > '9' {
			return false
		}
	}

	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}

	return true
}

func checkExpNum(parts []string) bool {
	if _, err := strconv.ParseFloat(parts[0], 64); err != nil {
		return false
	}

	if _, err := strconv.ParseInt(parts[1], 10, 64); err != nil {
		return false
	}
	return true
}
