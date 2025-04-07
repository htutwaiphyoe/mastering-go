package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	weekdays := days[:5]
	fmt.Println(weekdays)

	weekends := days[5:]
	fmt.Println(weekends)

	birthday := weekends[0]
	fmt.Println(birthday)

	advancedSE := days[3:5]
	fmt.Println(advancedSE)

	TGIF := advancedSE[1:]
	fmt.Println(TGIF)

	var productNames [5]string
	productNames[0] = "MacBook Pro M4Pro"
	productNames[1] = "iPhone 16 Pro"
	fmt.Println(productNames)

	ages := []int{18, 20, 22, 24, 26}

	adults := ages[1:]

	fmt.Println(adults)

	adults[0] = 21

	fmt.Println(adults)
	fmt.Println(ages)
	fmt.Println(len(adults))
	fmt.Println(cap(adults))

	adults = append(adults, 30)
	fmt.Println(adults)

	exercise()
}
