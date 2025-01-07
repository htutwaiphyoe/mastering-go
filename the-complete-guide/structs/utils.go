package main

import "fmt"

func getInput(prompt string) (value string) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
