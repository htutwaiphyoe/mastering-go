package firstClass

import "fmt"

type fn func(float64) float64

func FirstClass() {

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

func multiply(by float64) fn {

	fn := func(x float64) float64 {
		return x * by
	}

	return fn
}
