package main

import (
	"fmt"
	"strings"
)

type Email string

func (email Email) Validate() {
	if strings.Contains(string(email), "@") && strings.Contains(string(email), ".") {
		fmt.Println("Email is valid")

	} else {
		fmt.Println("Email is invalid")
	}
}

func main() {
	var email Email = "asdfadf.com"
	email.Validate()
}
