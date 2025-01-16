package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	weekdays := days[:5]
	fmt.Println(weekdays)

	weekends := days[5:]
	fmt.Println(weekends)

	birthday := weekends[:1]
	fmt.Println(birthday)
}
