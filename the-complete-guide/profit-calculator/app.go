package main

import "fmt"

func main() {
	var revenue, expenses, taxPercent float64

	revenue = getValue("Enter the revenue: ")
	expenses = getValue("Enter the expenses: ")
	taxPercent = getValue("Enter the tax rate: ")

	earningsBeforeTax, earningsAfterTax := calculateEarnings(revenue, expenses, taxPercent)

	ebt := fmt.Sprintf("Earnings before tax: %.2f\n", earningsBeforeTax)
	eat := fmt.Sprintf("Earnings after tax: %.2f\n", earningsAfterTax)

	fmt.Println(ebt, eat)
}

func getValue(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func calculateEarnings(revenue, expenses, taxPercent float64) (earningsBeforeTax float64, earningsAfterTax float64) {
	earningsBeforeTax = revenue - expenses
	earningsAfterTax = earningsBeforeTax * (1 - taxPercent/100)

	return earningsBeforeTax, earningsAfterTax
}
