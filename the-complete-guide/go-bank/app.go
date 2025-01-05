package main

import "github.com/htutwaiphyoe/mastering-go/the-complete-guide/go-bank/utils"

const file = "balance.txt"

func main() {
	var balance, err = utils.ReadFromFile(file)
	if err != nil {
		utils.SaveToFile(balance, file)
	}

	greet()

	for {
		action := start()

		switch action {
		case 1:
			check(balance)
		case 2:
			balance = deposit(balance)
		case 3:
			balance = withdraw(balance)
		case 4:
			exit()
			return
		default:
			invalid()
		}
	}
}
