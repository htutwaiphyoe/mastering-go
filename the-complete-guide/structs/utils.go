package main

import "fmt"

func getInput(prompt string) (value string) {
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}
