package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// InputData represents the input JSON structure
type InputData map[string][]float64

// OutputItem represents an item in the output JSON
type OutputItem struct {
	Tax   float64   `json:"tax"`
	Price []float64 `json:"price"`
	Total []float64 `json:"total"`
}

func main() {
	// Read input JSON file
	inputFile := "prices.json"
	outputFile := "prices-with-tax.json"

	// Read the JSON file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the JSON data
	var inputData InputData
	if err := json.Unmarshal(data, &inputData); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Process the data
	var outputItems []OutputItem
	for taxStr, prices := range inputData {
		// Convert tax string to float (remove the leading dot)
		taxRate, err := strconv.ParseFloat(taxStr, 64)
		if err != nil {
			// Try with the leading dot
			taxRate, err = strconv.ParseFloat(taxStr, 64)
			if err != nil {
				fmt.Printf("Error parsing tax rate '%s': %v\n", taxStr, err)
				continue
			}
		}

		// Calculate totals
		totals := make([]float64, len(prices))
		for i, price := range prices {
			totals[i] = price * (1 + taxRate)
		}

		// Create output item
		item := OutputItem{
			Tax:   taxRate,
			Price: prices,
			Total: totals,
		}
		outputItems = append(outputItems, item)
	}

	// Write output JSON
	outputJSON, err := json.MarshalIndent(outputItems, "", "  ")
	if err != nil {
		fmt.Printf("Error creating JSON: %v\n", err)
		return
	}

	// Write to file
	if err := os.WriteFile(outputFile, outputJSON, 0644); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Price calculation complete. Results written to %s\n", outputFile)
}
