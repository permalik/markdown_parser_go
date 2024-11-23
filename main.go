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
	f, e := os.Open("/Users/tymalik/Documents/git/markdown_parser_go/test.md")
	check(e)

	r := bufio.NewReader(f)
	lc := 0
	lex(r, lc)
}

func lex(r *bufio.Reader, lc int) {
	for {
		var l []byte
		for {
			d, i, e := r.ReadLine()
			if e != nil {
				if e.Error() == "EOF" {
					break
				}
				fmt.Println("error reading file: ", e)
				return
			}

			l = append(l, d...)

			if !i {
				break
			}
		}
		if len(l) == 0 {
			break
		} else {
			lc += 1
		}
		fmt.Printf("line#:%d\n%s\n", lc, string(l))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
