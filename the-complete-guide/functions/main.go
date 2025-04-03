package main

import "fmt"

type fn func(float64) float64

func main() {

	number := []float64{1, 2.5, 3.0, 4}

	doubled := transform(&number, multiply(2))

	fmt.Println("Doubled:", doubled)

	tripled := transform(&number, multiply(3))

	fmt.Println("Tripled:", tripled)

	quadrupled := transform(&number, multiply(4))

	fmt.Println("Quadrupled:", quadrupled)

}

func transform(numbers *[]float64, fn fn) []float64 {
	result := make([]float64, len(*numbers))

	for i, v := range *numbers {
		result[i] = fn(v)
	}

	return result
}

func double[T int | float64](x T) T {
	return x * 2
}

func triple[T int | float64](x T) T {
	return x * 3
}

func multiply(by int) fn {
	if by == 2 {
		return double
	}
	if by == 3 {
		return triple
	}

	return func(x float64) float64 {
		return x * float64(by)
	}
}
