package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regEx = regexp.MustCompile(`(?i)< *a *href *= *"`)
var regExCloseA = regexp.MustCompile(`(?i)< */ *a *>`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var buffer bytes.Buffer
	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
		buffer.WriteString(" ")
	}
	var all_doc string = buffer.String()

	all_groups := regEx.FindAllStringSubmatchIndex(all_doc, -1)

	for i := range all_groups {
		group := all_groups[i]
		var offset_quote int = strings.Index(all_doc[group[1]:], `"`) + group[1] // warn it can return -1
		var url string = all_doc[group[1]:offset_quote]
		var atag_angclose int = strings.Index(all_doc[offset_quote:], `>`) + offset_quote

		content_inds := regExCloseA.FindStringIndex(all_doc[atag_angclose:])
		var acont string = all_doc[atag_angclose+1 : atag_angclose+content_inds[0]]

		var ix_left_angclose int = strings.Index(acont, "</")
		if ix_left_angclose >= 0 {
			var k int
			for k = ix_left_angclose - 1; k > 0; k-- {
				var current string = acont[k : k+1]
				if current == ">" {
					break
				}
			}
			acont = acont[k+1 : ix_left_angclose]
		} else {
			if strings.Index(acont, "/>") >= 0 {
				acont = ""
			}
		}
		fmt.Printf("%s,%s\n", strings.TrimSpace(url), strings.TrimSpace(acont))
	}
}
