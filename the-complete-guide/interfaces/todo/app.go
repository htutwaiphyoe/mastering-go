package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mastering-go/the-complete-guide/interfaces/utils"
)

type Todo struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return nil, errors.New("content should not be empty")
	}

	return &Todo{
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("\nTodo: %v\n\n", todo.Content)
}

func (todo Todo) Save() error {
	fileName := fmt.Sprintf("todo-%v.json", strings.ToLower(strings.ReplaceAll(todo.Content, " ", "-")))

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func Get() string {
	content := utils.Input("Content:")
	return content
}
