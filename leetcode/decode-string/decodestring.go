package decodestring

import (
	"container/list"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	stack := list.New()
	type item struct {
		num int
		str string
	}

	for ix := 0; ix < len(s); {
		ch := s[ix]
		switch {
		case ch >= '0' && ch <= '9':
			var num string
			for ; ix < len(s); ix++ {
				if s[ix] == '[' {
					ix++ // pass '[' character
					break
				}
				num += string(s[ix])
			}
			n, _ := strconv.Atoi(num)
			// println("push number: ", n)
			stack.PushFront(&item{num: n})
		case ch == ']':
			ix++
			var str string
			var it *item
			for {
				it = stack.Remove(stack.Front()).(*item)
				if it.str == "" {
					// println("pop from stack number: ", it.num)
					break
				}
				// println("pop from stack string: ", it.str)
				str = it.str + str
			}
			stack.PushFront(&item{str: strings.Repeat(str, it.num)})
			// println("push repeated string: ", strings.Repeat(str, it.num))
		default: // letter
			var l string
			for ; ix < len(s) && unicode.IsLetter(rune(s[ix])); ix++ {
				l += string(s[ix])
			}
			stack.PushFront(&item{str: l})
			// println("push string: ", l)
		}
	}

	res := make([]string, stack.Len())
	for ix := stack.Len() - 1; ix >= 0; ix-- {
		it := stack.Remove(stack.Front()).(*item)
		res[ix] = it.str
	}
	return strings.Join(res, "")
}
