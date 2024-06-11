package main

import (
	"fmt"
	"math"
)

func main() {
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
