package closure

func Closure() {
	countOne := counter()
	countTwo := counter()

	countOne()
	countOne()

	countTwo()
}

func counter() func() int {
	counter := 0
	increment := func() int {
		counter += 1
		return counter
	}

	return increment
}
