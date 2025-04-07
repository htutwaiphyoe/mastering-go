package main

import (
	"fmt"
)

func exercise() {
	hobbies := [3]string{"Coding", "Writing", "Gaming"}

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	firstTwoHobbies := hobbies[0:2]
	fmt.Println(firstTwoHobbies)

	lastTwoHobbies := firstTwoHobbies[1:3]
	fmt.Println(lastTwoHobbies)

	courseGoals := []string{"Learn Go", "Build a web app"}

	courseGoals[1] = "Build a full stack app with React"

	courseGoals = append(courseGoals, "Projects with Go")
	fmt.Println(courseGoals)

	type Product struct {
		id    int
		title string
	}

	products := []Product{
		{1, "AirPods Pro"},
		{2, "iPhone 16 Pro"}}

	products = append(products, Product{3, "MacBook Pro M4Pro"})

	fmt.Println(products)

}
