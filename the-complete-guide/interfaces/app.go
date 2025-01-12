package main

import (
	"fmt"

	"github.com/mastering-go/the-complete-guide/interfaces/note"
	"github.com/mastering-go/the-complete-guide/interfaces/todo"
	"github.com/mastering-go/the-complete-guide/interfaces/utils"
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
	err = utils.Save(note)

	if err != nil {
		return
	}

	todo.Display()
	err = utils.Save(todo)

	if err != nil {
		return
	}

}
