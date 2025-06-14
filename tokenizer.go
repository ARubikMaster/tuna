package main

type line struct {
	text string
}

func (l line) Tokenize() ([]string, error) {
	line_content := []string{}
	line_content = append(line_content, l.text)
	return line_content, nil
}

func tokenizer(text string) ([]string, error) {
	return nil, nil
}
