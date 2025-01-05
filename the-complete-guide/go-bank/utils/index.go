package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Input(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)

	return value
}

func SaveToFile(value float64, filePath string) {
	err := os.WriteFile(filePath, []byte(fmt.Sprint(value)), 0644)
	if err != nil {
		fmt.Println("Something went wrong in saving value!")
	}
}

func ReadFromFile(filePath string) (float64, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0.00, errors.New("failed to read the file")
	}

	result, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0.0, errors.New("failed to read the data")
	}
	return result, nil
}
