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

func ensureBrackets(input string) string {
    // Trim leading and trailing whitespace
    trimmed := strings.TrimSpace(input)

    // Check if the string starts with '[' and ends with ']'
    if len(trimmed) >= 2 && trimmed[0] == '[' && trimmed[len(trimmed)-1] == ']' {
        return trimmed
    }

    // If not, add the brackets
    return "[" + trimmed + "]"
}


func CleanAndFormatJSON(input string) (string, error) {

	input = strings.Trim(input, " \n")
	input = strings.ReplaceAll(input, "\n", " ")

	input = ensureBrackets(input)	

	return input, nil
}

