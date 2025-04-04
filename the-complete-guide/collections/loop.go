package main

func loop() {
	ages := []int{29, 30}

	for index, value := range ages {
		println(index, value)
	}

	ageMaps := map[string]int{
		"john": 29,
		"doe":  30,
	}

	for key, value := range ageMaps {
		println(key, value)
	}
}
