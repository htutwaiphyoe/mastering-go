package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(prompt string) (value string) {
	fmt.Printf("%s ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func getNote() (string, string) {
	title := input("Title:")
	content := input("Content:")
	return title, content
}
