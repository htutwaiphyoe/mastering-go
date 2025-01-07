package main

import (
	"time"
)

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

	user.get()
	user.clear()
	user.get()
}
