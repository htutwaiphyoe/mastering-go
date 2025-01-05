package main

import "fmt"

func greet() {
	fmt.Println("Welcome to Go Bank!")
}

func start() (action float64) {
	fmt.Println("----------------------------------------")
	fmt.Println("What action would you like to perform?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("----------------------------------------")

	action = input("Enter the number of your action: ")

	return action
}

func check(balance float64) {
	fmt.Printf("Your current balance is %.2f\n", balance)
}

func deposit(balance float64) float64 {
	for {
		amount := input("Enter the amount to deposit: ")

		if amount < 0 {
			fmt.Println("Invalid amount. Please try again!")
			continue
		}

		balance += amount
		save(balance)
		fmt.Printf("Your updated balance is %.2f\n", balance)
		return balance
	}
}

func withdraw(balance float64) float64 {
	for {
		amount := input("Enter the amount to withdraw: ")

		if amount < 0 || amount > balance {
			fmt.Println("Invalid amount. Please try again!")
			continue
		}

		balance -= amount
		save(balance)
		fmt.Printf("Your updated balance is %.2f\n", balance)
		return balance
	}
}

func exit() {
	fmt.Println("Thanks for using GO Bank services.")
}

func invalid() {
	fmt.Println("Invalid action. Please try again!")
}
