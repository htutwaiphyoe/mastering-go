package main

import "fmt"

func main() {
	var revenue, expenses, taxPercent float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxPercent)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxPercent/100)

	ebt := fmt.Sprintf("Earnings before tax: %.2f\n", earningsBeforeTax)
	eat := fmt.Sprintf("Earnings after tax: %.2f\n", earningsAfterTax)

	fmt.Println(ebt, eat)
}
