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
			fmt.Printf("BlankLine: ''\nLine: %d\n", l.line)
		}
		return Token{
			Literal: literal.BlankLine{},
			Line:    l.line,
		}, nil

	case strings.HasPrefix(line, "# "):
		if l.debug {
			fmt.Printf("HeadingOne: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingOne{
				Text: line,
			},
			Line: l.line,
		}, nil

	case strings.Compare(line, "---") == 0:
		if l.debug {
			fmt.Printf("HorizontalRuleHyphen: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HorizontalRuleHyphen{
				Text: line,
			},
			Line: l.line,
		}, nil

	case strings.HasPrefix(line, "- [ ] "):
		if l.debug {
			fmt.Printf("TaskListItem: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.TaskList{
				Text: strings.TrimPrefix(line, "- [ ]"),
			},
			Line: l.line,
		}, nil

	case strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "+ "):
		if l.debug {
			fmt.Printf("ListItem: %s\nLine: %d\n", line, l.line)
		}

		if strings.HasPrefix(line, "* ") {
			return Token{
				Literal: literal.ListItem{
					Text: strings.TrimPrefix(line, "* "),
				},
				Line: l.line,
			}, nil
		}

		if strings.HasPrefix(line, "+ ") {
			return Token{
				Literal: literal.ListItem{
					Text: strings.TrimPrefix(line, "+ "),
				},
				Line: l.line,
			}, nil
		}

		return Token{
			Literal: literal.ListItem{
				Text: strings.TrimPrefix(line, "- "),
			},
			Line: l.line,
		}, nil

	default:
		// if l.debug {
		// 	fmt.Printf("Paragraph: %s\nLine: %d\n", line, l.line)
		// }
		return Token{
			Literal: literal.Paragraph{
				Text: line,
			},
			Line: l.line,
		}, nil
	}
}
