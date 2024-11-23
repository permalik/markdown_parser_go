package main

import (
	"bufio"
	"fmt"
	"os"
)

type token struct {
	name        string
	kind        string
	value       string
	line_number int32
}

func main() {
	f, e := os.Open("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	check(e)

	r := bufio.NewReader(f)
	var d []byte

	d, e = r.ReadBytes('\n')
	check(e)

	lex(d)
}

func lex(l []byte) {
	fmt.Println(string(l))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
