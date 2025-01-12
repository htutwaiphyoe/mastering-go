package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
