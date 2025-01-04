package main

import "fmt"

func main() {
	var balance = 0.0
	fmt.Println("Welcome to Go Bank!")

	for {
		action := start()

		if action == 1 {
			fmt.Printf("Your current balance is %.2f\n", balance)
		} else if action == 2 {

			for {
				amount := getInput("Enter the amount to deposit: ")

				if amount < 0 {
					fmt.Println("Invalid amount. Please try again!")
					continue
				}

				balance += amount
				fmt.Printf("Your updated balance is %.2f\n", balance)
				break
			}

		} else if action == 3 {
			for {
				amount := getInput("Enter the amount to withdraw: ")

				if amount < 0 || amount > balance {
					fmt.Println("Invalid amount. Please try again!")
					continue
				}

				balance -= amount
				fmt.Printf("Your updated balance is %.2f\n", balance)
				break
			}

		} else if action == 4 {
			fmt.Println("Thanks for using GO Bank services.")
			break
		} else {
			fmt.Println("Invalid action. Please try again!")
		}
	}
}

func start() (action float64) {
	fmt.Println("----------------------------------------")
	fmt.Println("What action would you like to perform?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("----------------------------------------")

	action = getInput("Enter the number of your action: ")

	return action
}

func getInput(prompt string) (value float64) {
	fmt.Print(prompt)
	fmt.Scan(&value)

	return value
}
