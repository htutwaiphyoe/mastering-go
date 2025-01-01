package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount, expectedReturnRate, noOfYears float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&noOfYears)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), noOfYears)
	actualValue := futureValue / math.Pow((1+inflationRate/100), noOfYears)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Actual value: %.2f\n", actualValue)
}