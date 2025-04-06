package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount := getInput("Enter the investment amount: ")
	expectedReturnRate := getInput("Enter the expected return rate: ")
	noOfYears := getInput("Enter the number of years: ")

	futureValue, actualValue := calculateFutureValues(investmentAmount, expectedReturnRate, noOfYears)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Actual value: %.2f\n", actualValue)
}

func getInput(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}

func calculateFutureValues(investmentAmount, expectedReturnRate, noOfYears float64) (futureValue float64, actualValue float64) {
	const inflationRate = 2.5

	futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), noOfYears)
	actualValue = futureValue / math.Pow((1+inflationRate/100), noOfYears)

	return futureValue, actualValue
}
