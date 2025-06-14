package main

import (
	"strings"
)

// line is a struct that represents a line of code
type line struct {
	text string
}

// Tokenize splits a line of code into tokens
func (l line) Tokenize() ([]string, error) {
	line_content := []string{}
	line_content = append(line_content, l.text)
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
