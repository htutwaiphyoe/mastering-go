package main

import (
	"fmt"

	"github.com/masering-go/the-complete-guide/functions/closure"
	firstClass "github.com/masering-go/the-complete-guide/functions/first-class"
	"github.com/masering-go/the-complete-guide/functions/recursion"
	"github.com/masering-go/the-complete-guide/functions/variadic"
)

func main() {
	firstClass.FirstClass()

	closure.Closure()

	recursion.TowerOfHanoi(4, "A", "B", "C")

	rest := variadic.Sum(1, 2, 3, 4, 5)
	fmt.Println("rest:", rest)

	numbers := []float64{1, 2, 3, 4, 5}
	spread := variadic.Sum(numbers...)
	fmt.Println("spread:", spread)

}
