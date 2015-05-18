package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var regEx = regexp.MustCompile(`/\*`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line string = strings.TrimSpace(scanner.Text())
		var ixslashslash int = strings.Index(line, "//")
		if ixslashslash >= 0 {
			slashcomments = append(slashcomments, ixslashslash+count)
		}

	}

}
