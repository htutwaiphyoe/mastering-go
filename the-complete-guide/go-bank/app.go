package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	balance, action := start()

	fmt.Println(balance, action)

}

func start() (balance float64, action float64) {
	fmt.Println("What action would you like to perform?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	balance = 0.0
	action = getInput("Enter the number of your action: ")

	return balance, action
}

func getInput(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)
	return value
}
