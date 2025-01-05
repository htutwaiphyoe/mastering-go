package main

import (
	"fmt"
	"time"
)

func main() {
	age := 25
	birthYear := getBirthYear(&age)
	fmt.Printf("Birth year: %v\n", birthYear)
}

func getBirthYear(age *int) int {
	return time.Now().Year() - *age
}
