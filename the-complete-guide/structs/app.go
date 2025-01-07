package main

import "fmt"

func main() {
	firstName := getInput("Please enter your first name: ")
	lastName := getInput("Please enter your last name: ")
	birthDate := getInput("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := newUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.get()
	user.clear()
	user.get()
}
