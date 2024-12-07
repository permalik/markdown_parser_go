package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/permalik/markdown_parser_go/lex"
	"github.com/permalik/markdown_parser_go/literal"
)

func main() {
	var input io.Reader
	var output io.Writer

	input_file := flag.String("i", "", "input file (default: stdin)")
	output_file := flag.String("o", "", "output file (default: stdout)")
	format_flag := flag.String("f", "md", "output format (md | html)")
	flag.Parse()

	if *input_file != "" {
		file, err := os.Open(*input_file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	if *output_file != "" {
		file, err := os.Create(*output_file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	lexer := lex.NewLexer(input, true)
	parser := NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing markdown: %v\n", err)
		os.Exit(1)
	}

	var visitor Visitor
	switch *format_flag {
	case "md":
		visitor = NewMDGen(output)
	default:
		fmt.Fprintf(os.Stderr, "unknown format: %s\n", *format_flag)
		os.Exit(1)
	}

	ast.Accept(visitor)
}

type Parser struct {
	lexer *lex.Lexer
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

func NewParser(lexer *lex.Lexer) *Parser {
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
