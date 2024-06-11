package main

import (
	"fmt"
	"math"
)

func main()  {
	var investmentAmount = 10000
	var expectedReturnRate = 5.5
	var noOfYears = 10

	var futureValue = float64(investmentAmount) *  math.Pow((1 + expectedReturnRate / 100), float64(noOfYears))

	fmt.Println(futureValue)
}