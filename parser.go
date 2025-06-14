package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parser(tokens [][]string) error {
	for _, line := range tokens {
		if line[0] == "print" {
			if len(line) < 2 {
				return fmt.Errorf("error: missing argument for print statement")
			}

			if strings.HasPrefix(line[1], "\"") && strings.HasSuffix(line[1], "\"") {
				fmt.Println(line[1][1 : len(line[1])-1])
			} else if val, err := strconv.Atoi(line[1]); err == nil {
				fmt.Println(val)
			} else if val, err := strconv.ParseFloat(line[1], 64); err == nil {
				fmt.Println(val)
			}
		}
	}

	return nil
}
