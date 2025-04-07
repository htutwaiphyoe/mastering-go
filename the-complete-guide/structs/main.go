package main

import (
	"fmt"

	"github.com/mastering-go/the-complete-guide/structs/admin"
	"github.com/mastering-go/the-complete-guide/structs/user"
)

func main() {
	firstName := getInput("Please enter your first name: ")
	lastName := getInput("Please enter your last name: ")
	birthDate := getInput("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	user.Display()
	user.Clear()
	user.Display()

	admin, err := admin.New("admin@gmail.com", "12345", user)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin.Get()

}
