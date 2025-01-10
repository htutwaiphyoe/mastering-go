package main

import (
	"fmt"

	"github.com/mastering-go/the-complete-guide/notey/note"
)

func main() {
	title, content := getNote()

	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()

	if err != nil {
		fmt.Println("Note cannot be saved.")
		return
	}

	fmt.Println("Note saved successfully.")

}
