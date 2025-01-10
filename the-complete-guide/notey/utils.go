package main

import (
	"errors"
	"fmt"
)

func input(prompt string) (value string, err error) {
	fmt.Println(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("empty")
	}

	return value, nil
}
