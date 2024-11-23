package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	check(err)

	reader := bufio.NewReader(f)
	data, err := reader.ReadBytes('\n')
	check(err)

	fmt.Println(string(data))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
