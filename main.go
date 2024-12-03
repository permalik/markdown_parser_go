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
	var byte_offset int
	Lex(reader, &line, &byte_offset)
}

func Lex(reader *bufio.Reader, line *string, byte_offset *int) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			byte_array := []byte{'<', ' ', 'l', 'i', 'n', 'e', 'b', 'r', 'e', 'a', 'k', ' ', '>'}
			*line = string(byte_array)
			fmt.Printf("lb: _, bo: %d\n", *byte_offset)
		} else {
			for _, v := range scanner.Bytes() {
				*byte_offset += 1
				fmt.Printf("ch: %s, bo: %d\n", string(v), *byte_offset)
			}
			// *line = string(scanner.Bytes())
		}
		// fmt.Println(*line)
	}
}
