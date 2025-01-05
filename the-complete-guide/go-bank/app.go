package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const file = "balance.txt"

func main() {
	var balance, err = read()
	if err != nil {
		save(balance)
	}

	greet()

	for {
		action := start()

		switch action {
		case 1:
			check(balance)
		case 2:
			balance = deposit(balance)
		case 3:
			balance = withdraw(balance)
		case 4:
			exit()
			return
		default:
			invalid()
		}
	}
}

func input(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)

	return value
}

func save(balance float64) {
	err := os.WriteFile(file, []byte(fmt.Sprint(balance)), 0644)
	if err != nil {
		fmt.Println("Something went wrong in saving balance!")
	}
}

func read() (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 0.00, errors.New("failed to read the data")
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0.0, errors.New("failed to read the balance")
	}
	return balance, nil
}
