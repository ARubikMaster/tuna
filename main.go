package main

import (
	"fmt"
	"os"
)

// Main function that tokenizes and parses a file given in the command line
func main() {
	// Check if a file was given
	if len(os.Args) < 2 {
		err := fmt.Errorf("error: no file given in command line")
		fmt.Println(err)
		return
	}

	// Get the file path
	file := os.Args[1]

	// Creates a byte slice to store the file content
	var slice_contents []byte

	// Reads file content and stores it in a byte slice if it exists as an entire path or in working directory

	// Checks if the file exists
	_, err := os.Stat(file)

	// If the file exists
	if err == nil {
		// Read the file
		slice_contents, err = os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else { // If the file doesn't exist
		// Get the current working directory
		local_dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Read the file
		slice_contents, err = os.ReadFile(local_dir + "\\" + file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Converts the byte slice to a string
	string_content := string(slice_contents)

	// Tokenize and parse the string
	tokens, err := tokenizer(string_content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = parser(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
}
