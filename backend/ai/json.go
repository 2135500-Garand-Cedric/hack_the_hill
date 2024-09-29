package ai

import (
	// "encoding/json"
	// "fmt"


	"strings"
)

type Task struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}

func CleanAndFormatJSON(input string) (string, error) {

	input = strings.Trim(input, " \n")
	input = strings.ReplaceAll(input, "\n", " ")
	return input, nil
}

