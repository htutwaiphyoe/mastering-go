package main

import "fmt"

func main() {

	number := []float64{1, 2.5, 3.0, 4}

	doubled := transform(&number, double)

	fmt.Println("Doubled:", doubled)

	tripled := transform(&number, triple)

	fmt.Println("Tripled:", tripled)

}

func transform(numbers *[]float64, fn func(float64) float64) []float64 {
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
