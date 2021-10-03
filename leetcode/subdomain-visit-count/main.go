package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Printf("%#v\n", subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {
	visits := make(map[string]int, 1)

	for _, dom := range cpdomains {
		parts := strings.Split(dom, " ")
		count, _ := strconv.Atoi(parts[0])
		fullDomain := parts[1]

		subDoms := strings.Split(fullDomain, ".")
		domain := subDoms[len(subDoms)-1]
		visits[domain] += count

		for ix := len(subDoms) - 2; ix >= 0; ix-- {
			domain = subDoms[ix] + "." + domain
			visits[domain] += count
		}
	}

	var res []string
	for domain, count := range visits {
		// res = append(res, fmt.Sprintf("%d %s", count, domain))
		res = append(res, strconv.Itoa(count)+" "+domain)
	}
	return res
}
