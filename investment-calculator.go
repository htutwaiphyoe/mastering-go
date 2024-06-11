package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, noOfYears, expectedReturnRate := 10000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), noOfYears)
	realValue := futureValue / math.Pow((1+inflationRate/100), noOfYears)

	fmt.Println(realValue)
}
