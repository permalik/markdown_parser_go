package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var input io.Reader
	var output io.Writer

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	if len(os.Args) > 2 {
		file, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	lexer := NewLexer(input)
}

type Literal interface {
	isLiteral()
}

type ListItem struct {
	Text string
}

type Paragraph struct {
	Text string
}

type BlankLine struct{}

func (l ListItem) isLiteral()  {}
func (p Paragraph) isLiteral() {}
func (b BlankLine) isLiteral() {}

type Token struct {
	Element Literal
	Line    int
	Column  int
}

type Lexer struct {
	scanner *bufio.Scanner
	line    int
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		scanner: bufio.NewScanner(reader),
		line:    0,
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
		return Token{
			Element: BlankLine{},
			Line:    l.line,
		}, nil

	default:
		return Token{
			Element: Paragraph{
				Text: line,
			},
			Line: l.line,
		}, nil
	}
}

type Parser struct {
	lexer *Lexer
}

type Node interface {
	Accept(v Visitor)
}

type TreeNode struct {
	Children []Node
}

type ListNode struct {
	Items []string
}

type ParagraphNode struct {
	Text string
}

type Visitor interface {
	VisitList(n *ListNode)
	VisitParagraph(n *ParagraphNode)
}

func (n *ListNode) Accept(v Visitor)      { v.VisitList(n) }
func (n *ParagraphNode) Accept(v Visitor) { v.VisitParagraph(n) }

func NewParser(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() (Node, error) {
	tree := &TreeNode{}
	var currList []string

	for {
		token, err := p.lexer.NextToken()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error parsing: %w", err)
		}

		switch tok := token.Element.(type) {
		case Paragraph:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &ParagraphNode{
				Text: tok.Text,
			})
		}
	}
}
