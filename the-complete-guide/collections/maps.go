package main

import "fmt"

func maps() {
	websites := map[string]string{
		"google":   "www.google.com",
		"facebook": "www.facebook.com",
		"twitter":  "www.twitter.com",
	}

	fmt.Print(websites)
}
