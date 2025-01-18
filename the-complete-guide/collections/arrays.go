package main

import "fmt"

func main() {
	var days []string = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println(days[5])
	days[5] = "Sat"
	fmt.Println(days[5])
}
