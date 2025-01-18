package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days[5])

	weekdays := days[0:5]
	fmt.Println(weekdays)

	weekends := days[5:]
	fmt.Println(weekends)
}
