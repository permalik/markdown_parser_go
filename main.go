package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/permalik/markdown_parser_go/lex"
	"github.com/permalik/markdown_parser_go/parse"
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
	parser := parse.NewParser(lexer)

	ast, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing markdown: %v\n", err)
		os.Exit(1)
	}

	var visitor parse.Visitor
	switch *format_flag {
	case "md":
		visitor = NewMDGen(output)
	default:
		fmt.Fprintf(os.Stderr, "unknown format: %s\n", *format_flag)
		os.Exit(1)
	}

	ast.Accept(visitor)
}

type MDGen struct {
	writer io.Writer
}

func NewMDGen(writer io.Writer) *MDGen {
	return &MDGen{writer: writer}
}

func (g *MDGen) VisitTree(n *parse.TreeNode) {
	for _, child := range n.Children {
		child.Accept(g)
	}
}

func (g *MDGen) VisitHorizontalRuleHyphen(n *parse.HorizontalRuleHyphenNode) {
	fmt.Fprintf(g.writer, "%s\n", n.Text)
}

func (g *MDGen) VisitList(n *parse.ListNode) {
	for _, item := range n.Items {
		fmt.Fprintf(g.writer, "* %s\n", item)
	}
	fmt.Fprintln(g.writer)
}

func (g *MDGen) VisitParagraph(n *parse.ParagraphNode) {
	fmt.Fprintf(g.writer, "%s\n\n", n.Text)
}
