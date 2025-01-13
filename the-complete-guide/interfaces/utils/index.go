package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

type Outputtable interface {
	Saver
	Display()
}

func Input(prompt string) (value string) {
	fmt.Printf("%s ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}

func Save(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving completed.")
	return nil
}

func Output(data Outputtable) error {
	data.Display()
	return Save(data)
}
