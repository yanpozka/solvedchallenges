package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regExp = regexp.MustCompile(`(?i)<[ ]*div[ ]*class[ ]*=[ ]*"[ ]*summary[ ]*"[ ]*>`)
var rExOpenH3 = regexp.MustCompile(`(?i)<[ ]*h3[ ]*>`)
var rExCloseH3 = regexp.MustCompile(`(?i)<[ ]*/[ ]*h3[ ]*>`)
var rExSpan = regexp.MustCompile(`(?i)asked[ ]*<[ ]*span[ ]+title[ ]*=[ ]*`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer

	for scanner.Scan() {
		buffer.WriteString(scanner.Text())
	}
	var all_doc string = buffer.String()
	indices := regExp.FindAllStringIndex(all_doc, -1)

	for i := range indices {
		// <div class="summary">
		var end int = indices[i][1]

		// <h3>
		h3 := rExOpenH3.FindStringIndex(all_doc[end:]) // if h3 == nil then Panic()
		var offset_h3 int = h3[1]

		close_h3 := rExCloseH3.FindStringIndex(all_doc[end+offset_h3:]) // if close_h3 == nil then Panic()
		var offset_close_h3 int = close_h3[0]

		var end_h3 int = end + offset_h3 + offset_close_h3
		var header string = all_doc[end+offset_h3 : end_h3]

		var title, id string = extractTitleAndId(header)
		fmt.Printf("%s;%s;", id, title)
		// </h3>
		// <span title=
		span := rExSpan.FindStringIndex(all_doc[end_h3:])
		var position_span_time int = span[1] + end_h3

		var start_time int = strings.Index(all_doc[position_span_time:], ">") + position_span_time + 1
		var end_time int = strings.Index(all_doc[position_span_time+1:], "<") + position_span_time + 1

		var time string = all_doc[start_time:end_time]
		fmt.Println(time)

		// </span>

	}
}

// TODO
func extractTitleAndId(hline string) (string, string) {
	return hline, "id1"
}
