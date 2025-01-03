package main

import "fmt"

func main() {
	revenue := getInput("Enter the revenue: ")
	expenses := getInput("Enter the expenses: ")
	taxPercent := getInput("Enter the tax rate: ")

	earningsBeforeTax, earningsAfterTax := calculateEarnings(revenue, expenses, taxPercent)

	ebt := fmt.Sprintf("Earnings before tax: %.2f\n", earningsBeforeTax)
	eat := fmt.Sprintf("Earnings after tax: %.2f\n", earningsAfterTax)

	fmt.Println(ebt, eat)
}

func getInput(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func calculateEarnings(revenue, expenses, taxPercent float64) (earningsBeforeTax float64, earningsAfterTax float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxPercent/100)

	return earningsBeforeTax, earningsAfterTax
}
