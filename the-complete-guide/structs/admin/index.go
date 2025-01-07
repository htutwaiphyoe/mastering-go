package admin

import (
	"errors"
	"fmt"

	"github.com/mastering-go/the-complete-guide/structs/user"
)

type Admin struct {
	email    string
	password string
	user.User
}

func (admin *Admin) Get() {
	fmt.Printf("Email: %s\n", admin.email)
	fmt.Printf("Password: %s\n", admin.password)
	admin.Display()
}

func New(email, password string, user *user.User) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("please provide all fields")
	}

	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}
