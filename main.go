package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/permalik/markdown_parser_go/literal"
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

	lexer := NewLexer(input, true)
	parser := NewParser(lexer)
	gen := NewMDGen(output)

	ast, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing markdown: %v\n", err)
		os.Exit(1)
	}

	ast.Accept(gen)
}

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

type Parser struct {
	lexer *Lexer
}

type Node interface {
	Accept(v Visitor)
}

type TreeNode struct {
	Children []Node
}

type HorizontalRuleHyphenNode struct {
	Text string
}

type ListNode struct {
	Items []string
}

type ParagraphNode struct {
	Text string
}

type Visitor interface {
	VisitTree(n *TreeNode)
	VisitHorizontalRuleHyphen(n *HorizontalRuleHyphenNode)
	VisitList(n *ListNode)
	VisitParagraph(n *ParagraphNode)
}

func (n *TreeNode) Accept(v Visitor)                 { v.VisitTree(n) }
func (n *HorizontalRuleHyphenNode) Accept(v Visitor) { v.VisitHorizontalRuleHyphen(n) }
func (n *ListNode) Accept(v Visitor)                 { v.VisitList(n) }
func (n *ParagraphNode) Accept(v Visitor)            { v.VisitParagraph(n) }

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

		switch tok := token.Literal.(type) {
		case literal.HorizontalRuleHyphen:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &HorizontalRuleHyphenNode{
				Text: tok.Text,
			})
		case literal.Paragraph:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
			tree.Children = append(tree.Children, &ParagraphNode{
				Text: tok.Text,
			})
		case literal.ListItem:
			currList = append(currList, tok.Text)
		case literal.BlankLine:
			if len(currList) > 0 {
				tree.Children = append(tree.Children, &ListNode{Items: currList})
				currList = nil
			}
		}
	}

	if len(currList) > 0 {
		tree.Children = append(tree.Children, &ListNode{Items: currList})
	}

	return tree, nil
}

type MDGen struct {
	writer io.Writer
}

func NewMDGen(writer io.Writer) *MDGen {
	return &MDGen{writer: writer}
}

func (g *MDGen) VisitTree(n *TreeNode) {
	for _, child := range n.Children {
		child.Accept(g)
	}
}

func (g *MDGen) VisitHorizontalRuleHyphen(n *HorizontalRuleHyphenNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
}

func (g *MDGen) VisitList(n *ListNode) {
	for _, item := range n.Items {
		fmt.Fprintf(g.writer, "* %s\n", item)
	}
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitParagraph(n *ParagraphNode) {
	fmt.Fprintf(g.writer, "%s\n\n", n.Text)
}
