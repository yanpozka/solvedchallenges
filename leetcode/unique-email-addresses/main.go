package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a.1+b+c@yan.com", "a.2+b+c@yan.com"}))
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}))
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}))
}

func numUniqueEmails(emails []string) int {
	uniqs := map[string]bool{}

	for _, email := range emails {
		ix := strings.Index(email, "@")
		user, domain := email[:ix], email[ix+1:]

		plusIx := strings.Index(user, "+")
		if plusIx == -1 {
			plusIx = len(user)
		}
		user = user[:plusIx]
		user = strings.ReplaceAll(user, ".", "")

		uniqs[user+"#"+domain] = true
	}

	return len(uniqs)
}

func numUniqueEmailsSplit(emails []string) int {
	uniqs := map[string]bool{}

	for _, email := range emails {
		parts := strings.Split(email, "@")
		user, domain := parts[0], parts[1]

		user = strings.Split(user, "+")[0]
		user = strings.ReplaceAll(user, ".", "")

		uniqs[user+"#"+domain] = true
	}

	return len(uniqs)
}
