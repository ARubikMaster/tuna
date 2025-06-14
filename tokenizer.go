package main

import (
	"slices"
	"strings"
)

var operators = []string{"+", "-", "*", "/", "%", "^", ">", "<", "=", "[", "]", "{", "}", "(", ")", ","}

// line is a struct that represents a line of code
type line struct {
	text string
}

// Tokenize splits a line of code into tokens
func (l line) Tokenize() ([]string, error) {
	l.text = strings.TrimRight(l.text, " \t\r")

	// Line content variable to be appended to
	line_content := []string{}

	// Variables to track the state of the tokenizer
	start_line := true
	var string_creating bool

	// Current token for tokens larger than 1 character
	var current_token string

	for _, char := range l.text {
		// Whitespace for indentation
		if char == ' ' && start_line {
			line_content = append(line_content, " ")

			continue
		} else {
			start_line = false
		}

		// Checks if a string is being created
		if char == '"' {
			string_creating = !string_creating
			current_token += "\""

			if !string_creating {
				line_content = append(line_content, current_token)
				current_token = ""
			}
			continue
		}

		// Checks if a string is being created
		if string_creating {
			current_token += string(char)
			continue
		}

		// Checks if the current character is an operator
		if slices.Contains(operators, string(char)) {
			if current_token != "" {
				line_content = append(line_content, current_token)
			}

			line_content = append(line_content, string(char))
			current_token = ""

			continue
		}

		// Checks if the current character is a space
		if char == ' ' {
			if current_token != "" {
				line_content = append(line_content, current_token)
			}
			current_token = ""

			continue
		}

		current_token += string(char)
	}

	// Appends the last token
	if current_token != "" {
		line_content = append(line_content, current_token)
	}

	//fmt.Println(strings.Join(line_content, " | "))

	return line_content, nil
}

// tokenizer splits a file of code into tokens
func tokenizer(text string) ([][]string, error) {
	// File tokens variable
	tokens := [][]string{}

	// Splits file content into lines
	lines_str := strings.Split(text, "\n")

	// Creates line structs for each line
	for _, line_str := range lines_str {
		line := line{line_str}

		// Tokenizes each line
		line_tokens, err := line.Tokenize()
		if err != nil {
			return nil, err
		}

		// Appends tokens to file tokens
		tokens = append(tokens, line_tokens)
	}

	return tokens, nil
}
