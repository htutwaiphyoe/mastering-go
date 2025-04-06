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

	err = utils.Output(note)

	if err != nil {
		return
	}

	err = utils.Output(todo)

	if err != nil {
		return
	}

}
