package gen

import (
	"fmt"
	"io"

	"github.com/permalik/markdown_parser_go/parse"
)

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
