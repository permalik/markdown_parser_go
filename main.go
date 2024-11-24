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

	ll := make([][]byte, 0)
	lc := 0
	readFile(r, &ll, lc)

	for i, v := range ll {
		fmt.Printf("line#:%d\n%s\n", i, string(v))
	}
}

func readFile(r *bufio.Reader, ll *[][]byte, lc int) {
	fmt.Println("log")
	s := bufio.NewScanner(r)
	for s.Scan() {
		lc += 1
		if len(s.Bytes()) == 0 {
			b := []byte{'<', ' ', 'n', 'e', 'w', 'l', 'i', 'n', 'e', ' ', '>'}
			*ll = append(*ll, b)
		} else {
			*ll = append(*ll, s.Bytes())
		}
	}
	// for {
	// 	var l []byte
	// 	for {
	// 		d, i, e := r.ReadLine()
	// 		if e != nil {
	// 			if e.Error() == "EOF" {
	// 				break
	// 			}
	// 			fmt.Println("error reading file: ", e)
	// 			return
	// 		}
	//
	// 		l = append(l, d...)
	//
	// 		if !i {
	// 			break
	// 		}
	// 	}
	// 	if len(l) == 0 {
	// 		n := []byte{'\n'}
	// 		*ll = append(*ll, n)
	// 	} else {
	// 		lc += 1
	// 		*ll = append(*ll, l)
	// 	}
	// }
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
