package main

import (
	"fmt"
	"time"
)

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	birthYear := getBirthYear(&age)
	fmt.Printf("Birth year: %v\n", birthYear)
}

func getBirthYear(age *int) int {
	return time.Now().Year() - *age
}
