package main

import (
	"fmt"
	"math"
)

func main()  {
	var investmentAmount float64 = 10000
	var expectedReturnRate float64 = 5.5
	var noOfYears  float64= 10

	var futureValue = investmentAmount *  math.Pow((1 + expectedReturnRate / 100), noOfYears)
	fmt.Println(futureValue)
}