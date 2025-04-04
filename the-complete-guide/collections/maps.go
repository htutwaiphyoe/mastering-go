package main

import "fmt"

type StringMap map[string][]string

func (stringMap StringMap) output() {
	for key, value := range stringMap {
		fmt.Println(key, value)
	}
}

func maps() {
	websites := map[string]string{
		"google":   "www.google.com",
		"facebook": "www.facebook.com",
		"twitter":  "www.twitter.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["google"])

	websites["linkedin"] = "www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "twitter")

	websites["x"] = "www.x.com"
	fmt.Println(websites)

	slices := make([]string, 0, 5)

	slices = append(slices, "a")
	slices = append(slices, "b")
	slices = append(slices, "c")

	fmt.Println(slices)

	maps := make(StringMap, 5)
	maps["a"] = []string{"a", "b", "c"}
	maps["b"] = []string{"d", "e", "f"}
	maps["c"] = []string{"g", "h", "i"}

	maps.output()
}
