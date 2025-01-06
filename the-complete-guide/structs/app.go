package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getInput("Please enter your first name: ")
	lastName := getInput("Please enter your last name: ")
	birthDate := getInput("Please enter your birthdate (MM/DD/YYYY): ")

	user := User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	fmt.Println(user)
}

func getInput(prompt string) (value string) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
