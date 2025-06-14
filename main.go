package main

import (
	"fmt"
	"log"
	"os"
)

// Main function that tokenizes and parses a file given in the command line
func main() {
	// Get the file path
	file := os.Args[1]

	// Reads file content and stores it in a byte slice if it exists as an entire path or in working directory
	// TODO: Fix this mess
	slice_contents_global, err_global := os.ReadFile(file)

	local_dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	slice_contents_local, err_local := os.ReadFile(local_dir + "/" + file)

	if err_global != nil && err_local != nil {
		err := fmt.Errorf("Error reading file")
		log.Fatal(err)
	}

	// Selects which slice to use and writes it to slice_contents
	var slice_contents []byte

	if err_local != nil {
		slice_contents = slice_contents_global
	} else {
		slice_contents = slice_contents_local
	}

	// Converts the byte slice to a string
	string_content := string(slice_contents)

	fmt.Println(string_content)

	// Tokenize and parse the string
	tokens, err := tokenizer(string_content)
	if err != nil {
		log.Fatal(err)
	}

	err = parser(tokens)
	if err != nil {
		log.Fatal(err)
	}
}
