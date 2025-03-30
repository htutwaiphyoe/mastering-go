package main

import "fmt"

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
}
