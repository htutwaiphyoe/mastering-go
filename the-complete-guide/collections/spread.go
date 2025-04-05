package main

import "fmt"

func spread() {
	users := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	newUsers := []string{"Frank", "Grace", "Heidi"}

	users = append(users, newUsers...)

	fmt.Println(users)
}
