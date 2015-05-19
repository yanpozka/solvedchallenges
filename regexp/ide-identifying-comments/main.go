package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regEx = regexp.MustCompile(`/\*`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var flag_is_commented = false
	var buffer bytes.Buffer

	for scanner.Scan() {
		var line string = strings.TrimSpace(scanner.Text())

		var ixslashslash int = strings.Index(line, "//")
		var ixasterix int = strings.Index(line, "/*")

		if flag_is_commented {
			var ixclosecomm int = strings.Index(line, "*/")
			if ixclosecomm >= 0 {
				buffer.WriteString(line[0 : ixclosecomm+2])

				fmt.Println(buffer.String())

				flag_is_commented = false
				buffer.Reset()

			} else {
				buffer.WriteString(line)
				buffer.WriteString("\n")
				continue
			}
		}

		if ixslashslash >= 0 && !flag_is_commented && (ixslashslash < ixasterix || ixasterix < 0) {
			fmt.Println(line[ixslashslash:])
			continue
		}

		if ixasterix >= 0 {
			buffer.WriteString(line[ixasterix:])
			buffer.WriteString("\n")
			flag_is_commented = true

			var ixclosecomm int = strings.Index(line, "*/")
			if ixclosecomm >= 0 {
				flag_is_commented = false
				fmt.Println(line[ixasterix : ixclosecomm+2])
				buffer.Reset()

				if ixslashslash >= 0 {
					fmt.Println(line[ixslashslash:])
				}
			}
		}
	}

}
