package main

import (
	"fmt"

	"github.com/mastering-go/the-complete-guide/interfaces/note"
	"github.com/mastering-go/the-complete-guide/interfaces/todo"
)

func main() {
	title, content := note.Get()
	todoInput := todo.Get()

	note, err := note.New(title, content)
	todo, todoErr := todo.New(todoInput)

	if err != nil || todoErr != nil {
		fmt.Println(err)
		return
	}

	note.Display()
	todo.Display()

	err = note.Save()
	todoErr = todo.Save()

	if err != nil {
		fmt.Println("Note cannot be saved.")
		return
	}

	if todoErr != nil {
		fmt.Println("Todo cannot be saved.")
		return
	}

	fmt.Println("Note saved successfully.")
	fmt.Println("Todo saved successfully.")

}
