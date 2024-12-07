package lex

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/permalik/markdown_parser_go/literal"
)

type Token struct {
	Literal literal.MarkdownLiteral
	Line    int
	Column  int
}

type Lexer struct {
	scanner *bufio.Scanner
	line    int
	debug   bool
}

func NewLexer(reader io.Reader, debug bool) *Lexer {
	return &Lexer{
		scanner: bufio.NewScanner(reader),
		debug:   debug,
	}
}

func (l *Lexer) NextToken() (Token, error) {
	if !l.scanner.Scan() {
		return Token{}, io.EOF
	}

	l.line++
	line := l.scanner.Text()

	switch {
	case len(line) == 0:
		if l.debug {
			fmt.Printf("BlankLine: '' Line: %d\n", l.line)
		}
		return Token{
			Literal: literal.BlankLine{},
			Line:    l.line,
		}, nil

	case strings.HasPrefix(line, "---"):
		return Token{
			Literal: literal.HorizontalRuleHyphen{
				Text: line,
			},
			Line: l.line,
		}, nil

	case strings.HasPrefix(line, "* "):
		return Token{
			Literal: literal.ListItem{
				Text: strings.TrimPrefix(line, "* "),
			},
			Line: l.line,
		}, nil

	default:
		return Token{
			Literal: literal.Paragraph{
				Text: line,
			},
			Line: l.line,
		}, nil
	}
}
