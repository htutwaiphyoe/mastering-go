package variadic

func Sum(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total
}
