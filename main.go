package main

import (
	"bufio"
	"fmt"
	"os"
)

type Token struct {
	Name       string
	Kind       string
	Value      string
	LineNumber int32
}

func main() {
	file, err := os.Open("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	if err != nil {
		fmt.Println("failed to open file")
		panic(err)
	}

	reader := bufio.NewReader(file)

	var line string
	Lex(reader, &line)
}

func Lex(reader *bufio.Reader, line *string) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			byte_array := []byte{'<', ' ', 'l', 'i', 'n', 'e', 'b', 'r', 'e', 'a', 'k', ' ', '>'}
			*line = string(byte_array)
		} else {
			*line = string(scanner.Bytes())
		}
		fmt.Println(*line)
	}
}
