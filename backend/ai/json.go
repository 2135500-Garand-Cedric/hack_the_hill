package ai

import (
	"encoding/json"
	// "fmt"


	"strings"
)

type Task struct {
	Task        string `json:"task"`
	Description string `json:"description"`
}

func CleanAndFormatJSON(input string) (string, error) {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(input, "\\\"", "\""), "\"\"", "\""), " ", ""), "\n", ""), "\\", ""), nil
}


func ConvertToValidJSON(input string) ([]Task, error) {
	var tasks []Task
	err := json.Unmarshal([]byte(input), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
