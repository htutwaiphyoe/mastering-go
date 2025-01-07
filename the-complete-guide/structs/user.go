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

func (user User) get() {
	fmt.Printf("First Name: %s\n", user.firstName)
	fmt.Printf("Last Name: %s\n", user.lastName)
	fmt.Printf("Birth Date: %s\n", user.birthDate)
}

func (user *User) clear() {
	user.firstName = ""
	user.lastName = ""
	user.birthDate = ""
	user.createdAt = time.Time{}
}
