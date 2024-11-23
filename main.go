package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	check(err)

	reader := bufio.NewReader(f)
	data := make([]byte, 100)
	_, err = reader.Read(data)
	check(err)

	fmt.Println(string(data))
}
