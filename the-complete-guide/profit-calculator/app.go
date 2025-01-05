package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getInput("Enter the revenue: ")

	if err != nil {
		fmt.Println("Revenue must not be less than zero")
		return
	}

	expenses, err := getInput("Enter the expenses: ")

	if err != nil {
		fmt.Println("Expense must not be less than zero")
		return
	}

	taxPercent, err := getInput("Enter the tax rate: ")

	if err != nil {
		fmt.Println("Tax rate must not be less than zero")
		return
	}

	earningsBeforeTax, earningsAfterTax := calculateEarnings(revenue, expenses, taxPercent)

	os.WriteFile("profile.txt", []byte(fmt.Sprintf("Ebt: %.2f\nEAT: %.2f\n", earningsBeforeTax, earningsAfterTax)), 0644)

	ebt := fmt.Sprintf("Earnings before tax: %.2f\n", earningsBeforeTax)
	eat := fmt.Sprintf("Earnings after tax: %.2f\n", earningsAfterTax)

	fmt.Println(ebt, eat)

}

func getInput(prompt string) (value float64, err error) {
	fmt.Print(prompt)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("invalid value")
	}
	return value, err
}

func calculateEarnings(revenue, expenses, taxPercent float64) (earningsBeforeTax float64, earningsAfterTax float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxPercent/100)

	return earningsBeforeTax, earningsAfterTax
}
