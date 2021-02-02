package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {

	uniqEmails := make(map[string]bool)
	for _, email := range emails {
		atIndex := strings.LastIndex(email, "@")

		plusIndex := strings.Index(email, "+")
		var username string
		if plusIndex != -1 {
			normalEmail := email[:plusIndex] + email[atIndex:]
			username = normalEmail[:plusIndex]
		} else {
			username = email[:atIndex]
		}

		// delete dot
		username = strings.ReplaceAll(username, ".", "")
		
		// join them back
		email = username + email[atIndex:]
		fmt.Println(email)
		uniqEmails[email] = true
	}

	return len(uniqEmails)

}

func main() {
	emails := []string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com"}

	result := numUniqueEmails(emails)

	expectedEmails := 2

	if result != expectedEmails {
		fmt.Println("something wrong")
	}

}
