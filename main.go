package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	check(err)
	fmt.Print(string(dat))
}
