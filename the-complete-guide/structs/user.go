package main

import (
	"errors"
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

func newUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("please provide all fields")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
