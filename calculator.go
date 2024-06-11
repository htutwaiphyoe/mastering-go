package main

import (
	"fmt"
	"math"
)

func calculateInvestment() {
	const inflationRate = 2.5
	var investmentAmount, noOfYears, expectedReturnRate float64

	fmt.Println("Enter investment amount:")
	fmt.Scan(&investmentAmount)

	fmt.Println("Enter expected return rate:")
	fmt.Scan(&expectedReturnRate)

	fmt.Println("Enter number of years:")
	fmt.Scan(&noOfYears)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), noOfYears)
	realValue := futureValue / math.Pow((1+inflationRate/100), noOfYears)

	fmt.Println("The total value: ", realValue)
}

func calculateProfit() {
	var revenue, expenses, tax float64

	fmt.Println("Enter revenue:")
	fmt.Scan(&revenue)

	fmt.Println("Enter expenses:")
	fmt.Scan(&expenses)

	fmt.Println("Enter tax:")
	fmt.Scan(&tax)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - (tax / 100))
	ratio := earningBeforeTax / earningAfterTax

	fmt.Println("Earning before tax: ", earningBeforeTax)
	fmt.Println("Earning after tax: ", earningAfterTax)
	fmt.Println("Ratio: ", ratio)
}

func main() {
	var calculation string

	fmt.Println("Enter calculation:")
	fmt.Scan(&calculation)

	if calculation == "profit" {
		calculateProfit()
	} else {
		calculateInvestment()
	}
}
