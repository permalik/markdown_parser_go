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
	case strings.HasPrefix(line, "## "):
		if l.debug {
			fmt.Printf("HeadingTwo: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingTwo{
				Text: line,
			},
			Line: l.line,
		}, nil
	case strings.HasPrefix(line, "### "):
		if l.debug {
			fmt.Printf("HeadingThree: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingThree{
				Text: line,
			},
			Line: l.line,
		}, nil
	case strings.HasPrefix(line, "#### "):
		if l.debug {
			fmt.Printf("HeadingFour: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingFour{
				Text: line,
			},
			Line: l.line,
		}, nil
	case strings.HasPrefix(line, "##### "):
		if l.debug {
			fmt.Printf("HeadingFive: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingFive{
				Text: line,
			},
			Line: l.line,
		}, nil
	case strings.HasPrefix(line, "###### "):
		if l.debug {
			fmt.Printf("HeadingSix: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HeadingSix{
				Text: line,
			},
			Line: l.line,
		}, nil
	// TODO: allow > 3 characters on line
	case strings.Compare(line, "---") == 0 || strings.Compare(line, "___") == 0 || strings.Compare(line, "***") == 0:
		if l.debug {
			fmt.Printf("HorizontalRuleHyphen: %s\nLine: %d\n", line, l.line)
		}
		return Token{
			Literal: literal.HorizontalRule{
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
	case strings.HasPrefix(line, ": "):
		return Token{
			Literal: literal.Definition{
				Text: strings.TrimPrefix(line, ": "),
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
