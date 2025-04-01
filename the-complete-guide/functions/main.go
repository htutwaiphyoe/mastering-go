package main

import (
	"fmt"

	"github.com/masering-go/the-complete-guide/functions/closure"
	"github.com/masering-go/the-complete-guide/functions/recursion"
	"github.com/masering-go/the-complete-guide/functions/variadic"
)

func main() {
	firstClass()

	closure.Closure()

	recursion.TowerOfHanoi(4, "A", "B", "C")

	sum := variadic.Sum(1, 2, 3, 4, 5)
	fmt.Println("variadic:", sum)

}
